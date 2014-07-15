// Package routes is the library for freedom-routes.
package routes

import (
	"fmt"
	"io/ioutil"
	"math"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"text/template"
)

const ASSETS_MODE = "source"

var ASSETS_PATH = getAssetsPath(ASSETS_MODE)
var REMOTE_URL = "http://ftp.apnic.net/apnic/stats/apnic/delegated-apnic-latest"
var LOCAL_PATH = "/etc/freedom-routes"

type Ip struct {
	Ip   string
	Cidr string
	Mask string
}

func (ip Ip) String() string {
	return fmt.Sprintf("%s/%s", ip.Ip, ip.Cidr)
}

// Generate result files from template.
func Generate(templateName string, ips []Ip, outputDir string) {
	templateDir := filepath.Join(ASSETS_PATH, "templates", templateName)

	type Data struct {
		Ips []Ip
	}
	data := Data{ips}

	files, err := ioutil.ReadDir(templateDir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		templatePath := filepath.Join(templateDir, file.Name())
		output := filepath.Join(outputDir, file.Name())

		tmpl, err := template.ParseFiles(templatePath)
		if err != nil {
			panic(err)
		}
		file, err := os.Create(output)
		if err != nil {
			panic(err)
		}
		defer file.Close()
		if runtime.GOOS != "windows" {
			err = file.Chmod(0755)
			if err != nil {
				panic(err)
			}
		}
		err = tmpl.Execute(file, data)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Create %s\n", output)
	}
}

// Fetch local ips and remote ips.
func FetchIps() (ips []Ip) {
	ips = FetchLocalIps()
	ips = append(ips, FetchRemoteIps()...)

	return ips
}

// Fetch remote ips from REMOTE_URL.
//
// syntax:
//
//   apnic|CN|ipv4|1.94.0.0|131072|20100806|allocated
//   ...
func FetchRemoteIps() (ips []Ip) {
	println("正在从apnic.net下载ip地址库, 这需要好几分钟, 请耐心等待...")

	resp, err := http.Get(REMOTE_URL)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	re, _ := regexp.Compile(`apnic\|CN\|ipv4\|([\d\.]+)\|(\d+)\|`)
	rows := re.FindAllSubmatch(body, -1)
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		ip := string(row[1])
		num_ip, _ := strconv.Atoi(string(row[2]))
		cidr := 32 - int(math.Log2(float64(num_ip)))
		cidrStr := strconv.Itoa(cidr)

		ips = append(ips, Ip{ip, cidrStr, cidr2mask(cidr)})
	}

	return ips
}

// Fetch local ips from LOCAL_PATH.
//
// syntax:
//
//   70.33.217.25/32
//   ...
func FetchLocalIps() (ips []Ip) {
	data, err := ioutil.ReadFile(LOCAL_PATH)
	if os.IsNotExist(err) {
		return ips
	} else if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.Trim(line, " \n\t")
		if line == "" {
			continue
		}
		tokens := strings.Split(line, "/")
		ip := tokens[0]
		cidr := tokens[1]
		cidrInt, _ := strconv.Atoi(cidr)
		ips = append(ips, Ip{ip, cidr, cidr2mask(cidrInt)})
	}

	return ips
}

// Convert CIDR to Mask.
// 
// Example:
//
//   cidr2mask(24)       #-> "255.255.255.0"
func cidr2mask(cidr int) string {
	mask := net.CIDRMask(cidr, 32)
	masks := []string{}
	for _, v := range mask {
		masks = append(masks, strconv.Itoa(int(v)))
	}

	return strings.Join(masks, ".")
}
