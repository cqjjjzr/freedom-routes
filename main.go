package main

import (
	"github.com/SaberSalv/freedom-routes/routes"
	"fmt"
	"github.com/ogier/pflag"
	"path/filepath"
	"os"
)

const VERSION = "1.1.0"

var USAGE = `
$ freedom-routes [options] <template ..>

OPTIONS:
	-o, --output="."                 # output directory
	-h, --help
	--version
`

func genRoutes(templateNames []string, outputDir string) {
	ips := routes.FetchIps()

	if len(templateNames) == 1 {
		os.MkdirAll(outputDir, 0755)
		routes.Generate(templateNames[0], ips, outputDir)
	} else {
		for _, templateName := range templateNames {
			output := filepath.Join(outputDir, templateName)
			os.MkdirAll(output, 0755)
			routes.Generate(templateName, ips, output)
		}
	}
}

func main() {
	pflag.Usage = func() {
		fmt.Fprintf(os.Stderr, USAGE)
	}

	var output = pflag.StringP("output", "o", ".", "output directory")
	var version = pflag.BoolP("version", "", false, "version")
	pflag.Parse()

	if *version {
		fmt.Println(VERSION)
	} else if pflag.NArg() > 0 {
		genRoutes(pflag.Args(), *output)
	} else {
		pflag.Usage()
	}
}
