## sword
sword是一套简单易用的Go语言业务框架，整体逻辑设计简洁，支持HTTP服务、分布式调度和和脚本任务等常用业务场景模式。

## Goals
让PHPer更加容易切换角色成为Gopher。




## Quick start

### Requirements
- Go version >= 1.12
- Global environment configure (Linux/Mac)  

```
export  GOPROXY=https://goproxy.cn
export  GO111MODULE=on
```

### Installation
```shell
go get github.com/moka-mrp/sword
sword new sword-demo -p /tmp  -m moye
```

### Build & Run
```shell
cd /tmp/sword-demo
go run main.go  api
```

### Test demo
```
curl "http://127.0.0.1:8080/hello"
```

## Documents

- [项目地址](https://github.com/qit-team/sword)
- [中文文档](https://github.com/qit-team/sword/wiki)
- [changelog](https://github.com/qit-team/sword/blob/master/CHANGLOG.md)

## Contributors

- Tinson Ho ([@tinson](https://github.com/hetiansu5))
- ACoderHIT ([@ACoderHIT](https://github.com/ACoderHIT))
- xiongwilee ([@xiongwilee](https://github.com/xiongwilee))
- KEL ([@deathkel](https://github.com/deathkel))
- peterwu




