bash脚本上传生成的routes-{up,down}.sh文件到云存储.

## 使用方法

```
$ bin/uploader upload s3 *.sh
```

## Amazon S3

安装s3cmd

**ArchLinux**: `$ pacman -S s3cmd`


## 新浪云存储

[安装scs-cli-tool](http://open.sinastorage.com/?c=doc&a=sdk#scs-cli-tool)

## ref

```
# /etc/cron.hourly/freedom-routes

  #!/bin/bash

  cd <project>
  freedom-routes android linux mac routeros windows -o tmp
  misc/uploader/bin/uploader upload s3 tmp/*
```
