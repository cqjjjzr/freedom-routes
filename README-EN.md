freedom-routes, Intenet Freedom - add local routes in vpn for accessing local websites using direct link
========================================================================================================

|                |                                                             |
|----------------|------------------------------------------------------       |
| Homepage:      | https://github.com/SaberSalv/freedom-routes        |
| Author:	       | Guten                                            |
| License:       | MIT-LICENSE                                                |
| Documentation: | http://godoc.org/github.com/SaberSalv/freedom-routes
| Issue Tracker: | https://github.com/SaberSalv/freedom-routes/issues |
| Platforms:     | Linux, Mac OS X, Windows                         |

Use ip data from APNIC to generate routes for vpn to make accessing websites in China use direct link, while accessing foreign websites use vpn. Which make browse websites much faster.

*Please Help*: the mac package can't get into homebrew's main repository for it's un-notable, so if you like the project, please star it at the top of this page.

Getting started
---------------

**Usage**:

	$ freedom-routes [options] <template>
      -o, --output="."                 # output directory

View all avaliable templates at [here](https://github.com/SaberSalv/freedom-routes/tree/master/routes/templates)

**Linux (OpenVPN)**

	$ freedom-routes linux -o /etc/openvpn
	> create /etc/openvpn/routes-up.sh
	> create /etc/openvpn/routes-down.sh

	$ edit /etc/openvpn/x.conf

		script-security 2
		up ./routes-up.sh
		down ./routes-down.sh

**GUI Application for Windows**
    Windows user can use GUI Application to auto update route table. Please install Java Runtime Environment to execute the application. Download:http://pan.baidu.com/s/1c0u9lte

**RouterOS**

	$ freedom-routes routeros

* upload freedomroutes.rsc via winbox/WebConsole/ssh/ftp
* get into your router terminal, type `/import freedomroutes.rsc`
* add a default route with `routing-mark` set to `freedomroutes.domestic`
* see `/ip route rule print` for more details..

Others, see documentation at [here](https://github.com/SaberSalv/chnroutes/blob/master/README.md) (Chinese)

**Add IPs from local**

	# create /etc/freedom-routes
		70.33.217.1/32

**Auto update IP data**

Ip data may change over time, so it's recommand to update every month.

	$ create /etc/cron.monthly/freedom-routes with 0755 mode

		#!/bin/bash

		freedom-routes linux -o /etc/openvpn

Install
-------

**Use Package**

* ArchLinux: `pacaur -S freedom-routes` # install [freedom-routes](https://aur.archlinux.org/packages/freedom-routes/) from AUR
* Mac OS X: `brew install sabersalv/alt/freedom-routes`
* Windows: download [x64](http://dl.saber.li/freedom-routes/freedom-routes.windows.amd64.zip) [x86](http://dl.saber.li/freedom-routes/freedom-routes.windows.386.zip)

**Use Source**

	$ [sudo] go get github.com/SaberSalv/freedom-routes
	
**Dependencies**

* iproute2 (Linux, Mac)

Development 
===========

Build
------

	mkdir output
	sed -i '/const ASSETS_MODE/s/.*/const ASSETS_MODE = "runtime"/' routes/routes.go
	go build -o output/freedom-routes
	cp -r routes/templates output

Contributing 
-------------

* Submit any bugs/features/ideas to github issue tracker.

Please see [Contibution Documentation](https://github.com/SaberSalv/freedom-routes/blob/master/CONTRIBUTING.md).

A list of [Contributors](https://github.com/SaberSalv/freedom-routes/contributors).

Resources
---------

* [chnroutes](https://github.com/SaberSalv/chnroutes): scripts help chinese netizen with vpn.
* [Pacaur](https://github.com/Spyhawk/pacaur): An AUR helper for ArchLinux.
* [Homebrew](https://github.com/mxcl/homebrew): The missing package manager for Mac OS X.

Copyright
---------

(the MIT License)

Copyright (c) 2013 Guten

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the 'Software'), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
