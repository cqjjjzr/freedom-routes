freedom-routes, chnroutes的改进版本, 大幅提升VPN浏览国内网页速度.
================================================================

|                |                                                             |
|----------------|------------------------------------------------------       |
| 主页:          | https://github.com/SaberSalv/freedom-routes        |
| 作者:	         | Saber                                            |
| 版权:          | MIT-LICENSE                                                |
| 提交Bug:       | https://github.com/SaberSalv/freedom-routes/issues   |
| API文档        | http://godoc.org/github.com/SaberSalv/freedom-routes |
| 支持平台:      | Linux, Mac OS X, Windows, OpenWRT, DD-WRT, ASUSWRT   |

生成一个可以运行的脚本, 当VPN运行的时候, 自动添加国内的IP地址到系统`路由表`, 用`直接连接`方式访问国内的网站, 用`VPN`方式访问国外的网站, 从而提升网页浏览速度. (例如: 使用前ping baidu.com是300ms延迟, 使用后可以减少到30ms)

## 对chnroutes的改进

1. Linux下导入路由的速度更快, 秒时间导入
2. 支持模板, 可以自定义脚本

# 网络版本

## 下载

每24小时更新一次.

- **Linux, OpenWRT, DD-WRT, ASUSWRT**: [routes-up.sh](http://dl.saber.li/freedom-routes/linux/routes-up.sh) [routes-down.sh](http://dl.saber.li/freedom-routes/linux/routes-down.sh)
- **Mac OS X**: [routes-up.sh](http://dl.saber.li/freedom-routes/mac/routes-up.sh) [routes-down.sh](http://dl.saber.li/freedom-routes/mac/routes-down.sh)
- **Windows**: [routes-up.bat](http://dl.saber.li/freedom-routes/windows/routes-up.bat) [routes-down.bat](http://dl.saber.li/freedom-routes/windows/routes-down.bat)
- **Android**: [routes-up.sh](http://dl.saber.li/freedom-routes/android/routes-up.sh) [routes-down.sh](http://dl.saber.li/freedom-routes/android/routes-down.sh)
- **RouterOS**: [freedomroutes.rsc](http://dl.saber.li/freedom-routes/routeros/freedomroutes.rsc)

## 使用方法

这些ip地址库并不是固定不变的, 尽管变化不大, 但还是建议每隔两三个月更新一次.

**手动运行**

```
启用
# ./routes-up.sh

停用
# ./routes-down.sh
```

Windows需要右键route-up.bat -> 已管理员身份运行. 并且导入需要好几分钟, 请耐心等待, 有哪位大神知道有更好的方法的话, 求指教.

**OpenVPN**

```
# cp them to /etc/openvpn
# edit /etc/openvpn/hello.conf

  script-security 2
  up ./routes-up.sh
  down ./routes-down.sh
```

**PPTP**

```
cp routes-up.sh /etc/ppp/ip-pre-up
cp routes-down.sh /etc/ppp/ip-down.d/ip-down
```

# 本地版本

查看[本地版本使用说明](https://github.com/SaberSalv/freedom-routes/blob/master/docs/local.md), 不过不推荐使用, 用网络版本更方便.

# 开发

## 任何人都可以帮助这个项目

- 请保持低调, 不要推广本项目, 自己用的可以就好了, 不要点击右上角的star按钮.
- 提交Bug/建议
- 帮助作者提高文档

感谢所有[贡献者](https://github.com/SaberSalv/freedom-routes/contributors). </br>
感谢原来的[chnroutes](https://github.com/fivesheep/chnroutes)作者.


## 编译

```
$ mkdir output
$ sed -i '/const ASSETS_MODE/s/.*/const ASSETS_MODE = "runtime"/' routes/routes.go
$ go build -o output/freedom-routes
$ cp -r routes/templates output
```

## 版权

(the MIT License)

Copyright (c) 2013-2015 Saber

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the 'Software'), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
