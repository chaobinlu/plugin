[![pipeline status](https://api.travis-ci.org/33cn/plugin.svg?branch=master)](https://travis-ci.org/33cn/plugin/)
[![Go Report Card](https://goreportcard.com/badge/github.com/33cn/plugin?branch=master)](https://goreportcard.com/report/github.com/33cn/plugin)

# chain33 官方插件系统

## 安装

##### 1. 安装govendor 工具

```
go get -u -v github.com/kardianos/govendor
```

```
C:\Users\pc005>go get -u -v github.com/kardianos/govendor
```
```
github.com/kardianos/govendor (download)
github.com/kardianos/govendor/vendor/github.com/Bowery/prompt
github.com/kardianos/govendor/prompt
github.com/kardianos/govendor/internal/pathos
github.com/kardianos/govendor/vendor/github.com/dchest/safefile
github.com/kardianos/govendor/cliprompt
github.com/kardianos/govendor/internal/vos
github.com/kardianos/govendor/pkgspec
github.com/kardianos/govendor/vendorfile
github.com/kardianos/govendor/vendor/github.com/pkg/errors
github.com/kardianos/govendor/internal/vfilepath
github.com/kardianos/govendor/vcs
github.com/kardianos/govendor/vendor/gopkg.in/yaml.v2
github.com/kardianos/govendor/vendor/github.com/google/shlex
github.com/kardianos/govendor/vendor/golang.org/x/tools/go/vcs
github.com/kardianos/govendor/context
github.com/kardianos/govendor/migrate
github.com/kardianos/govendor/help
github.com/kardianos/govendor/run
github.com/kardianos/govendor

```
#### 支持make file的平台
```

```
make
```
就可以完成编译安装

## 运行

```
./chain33 -f chain33.toml
```
注意，默认配置会连接chain33 测试网络

## 注意:

从头开始安装vendor 有非常大的难度，主要问题是带宽 和 翻墙问题
为了解决包依赖等问题，我们直接提供了vendor目录。


