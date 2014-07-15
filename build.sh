#!/bin/bash

# Zeus
#
#  ./build.sh release
#

# Build binary for distribution.
#
# Usage:
#
#   ./build.sh [options] [platform ...]
#     -d <assets_dir>
#     -p              # packaging the binary
#
# Example:
#
#   ./build.sh -d /usr/share/freedom-routes    # build with system go
#   ./build.sh release                         # build $RELEASEs with user complied go and upload the package to s3.

EXTRA_FILES="routes/templates freedom-routes.etc README.md"
RELEASE="linux/amd64 linux/386 homebrew/amd64 homebrew/386 windows/amd64 windows/386"
VERSION=$(sed -rn 's/.*const VERSION.*"([0-9.]+)".*/\1/p' main.go)
declare -A OS_MAP=(
	[homebrew]="darwin"
)
declare -A DIR_MAP=(
	[linux]="/usr/share/freedom-routes"
	[homebrew]="/usr/share/freedom-routes"
)

platform=""
os=""
arch=""
assets_dir=""

function cgo_enabled { [ "$1" == "$GOHOSTOS" ] && echo 1 || echo 0; }
function ext { [ $1 == "windows" ] && echo .exe || echo ""; }

# dist{platform, os, arch, assets_dir}
function dist {
	output="dist/$platform.$arch"
	rm -r $output 2>/dev/null
	mkdir -p $output

	cp -r $EXTRA_FILES $output
	if [[ -d "misc/$platform" ]]; then
		cp -r misc/$platform/* $output
	fi

	build
}

# build{platform, os, arch, assets_dir}
function build {
	echo -e "\n>> building $platform/$arch"
	output="dist/$platform.$arch/freedom-routes$(ext $os)"
	sed -i "/const ASSETS_MODE/s~.*~const ASSETS_MODE = \"$assets_dir\"~" routes/routes.go
	CGO_ENABLED=$(cgo_enabled $os) GOOS=$os GOARCH=$arch $GOROOT/bin/go build -o "$output"
	sed -i '/const ASSETS_MODE/s/.*/const ASSETS_MODE = "source"/' routes/routes.go
	echo ">> created $output"
}

# package{platform, os, arch}
function package {
	echo ">> packing $platform/$arch"

	case $os in
		linux | darwin )
			tar zcf dist/freedom-routes.$platform.$arch-$VERSION.tar.gz -C dist $platform.$arch;;
		windows )
			rm dist/freedom-routes.$platform.$arch-$VERSION.zip 2>/dev/null
			cd dist && zip -r freedom-routes.$platform.$arch-$VERSION.zip $platform.$arch && cd ..
			;;
	esac
}

function upload {
	for file in $*; do
		s3cmd put --acl-public $file s3://dl.saber.li/freedom-routes/
		#current=$(echo $file | sed -r 's/(.*)-[0-9.]+(\..*)$/\1\2/')
		#s3cmd copy s3://downloads.sabersalv.li/freedom-routes/$file s3://downloads.sabersalv.li/freedom-routes/$current
	done
}

#
# ¤main
# ----

eval "$(go env)"

o_assets_dir="runtime"
o_package=false
while getopts "d:p" opt; do
  case $opt in
    d ) o_assets_dir=$OPTARG ;;
		p ) o_package=true ;;
  esac
done
shift $(( OPTIND - 1 ))

case $1 in
	"" )
		platform=$GOOS; os=$GOOS; arch=$GOARCH; assets_dir=$o_assets_dir
		dist
		[ $o_package == true ] && package
		;;
	dist )
		for release in $RELEASE; do
			platform=${release%/*}; arch=${release#*/}; os=${OS_MAP[$platform]-$platform}
			assets_dir=${DIR_MAP[$platform]-$o_assets_dir}
			# fix linux/amd64 build error
			[[ $release == "linux/amd64" ]] && GOROOT="" || GOROOT="/home/guten/dev/src/go/go"
			dist
			package
		done
		;;
	release )
		rm dist/*.tar.gz dist/*.zip 2>/dev/null
		"$0" dist
		upload dist/*.tar.gz dist/*.zip
		;;
	upload )
		shift
		upload $*
		;;
	* )
		GOROOT="/home/guten/dev/src/go/go"
		platform=${1%/*}; arch=${1#*/}; os=${OS_MAP[$platform]-$platform}
		dist
		[ $o_package == true ] && package
		;;
esac

exit 0
