## 安装

- **ArchLinux**: 从AUR里面安装 `$ yaourt -S freedom-routes`
- **Mac OS X**: `$ brew install sabersalv/alt/freedom-routes`
- **Windows**: 下载[x64](https://github.com/SaberSalv/freedom-routes/releases/download/v1.1.0/freedom-routes.windows.amd64-1.1.0.zip)(64位), [x86](https://github.com/SaberSalv/freedom-routes/releases/download/v1.1.0/freedom-routes.windows.386-1.1.0.zip)(32位)

## 使用方法

**Usage**:

```
$ freedom-routes linux
> 生成routes-up.sh routes-down.sh
$ freedom-routes mac
$ freedom-routes windows

$ freedom-routes [options] <template>
    -o, --output="."                 # 输出脚本目录
```

查看所有的[模板](https://github.com/SaberSalv/freedom-routes/tree/master/routes/templates)

**本地的配置文件**

(Linux)

	# edit /etc/freedom-routes
		70.33.217.1/32

**自动更新**

(Linux)

	$ create /etc/cron.weekly/freedom-routes with 0755 mode

		#!/bin/bash

		freedom-routes linux -o /etc/openvpn
