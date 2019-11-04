# user服务模块结构
```$xslt
# 生成接口代码（proto）：
protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto

└─user
    │  Dockerfile    
    │  main.go         # user后端服务启动程序
    │  Makefile       
    │  plugin.go       # 插件 
    │  README.md
    ├─handler
    │      user.go     # 处理请求响应逻辑/逻辑实现
    ├─model            # 数据库初始化和具体的数据模型
    │      db.go       # 初始化数据库 `gorm`
    │      user.go     # user数据库模型
    ├─proto            # proto的接口实现/生成接口代码
    │  └─user
    │          user.micro.go
    │          user.pb.go
    │          user.proto
    ├─repository       # （crdu）数据处理的实现，增删改查
    │      user.go     # 处理user模块的增删改查
    └─subscriber
            user.go

```





# 以下为自动生成的README
# User Service

This is the User service

Generated with

```
micro new micros/user --namespace=go.micro --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.user
- Type: srv
- Alias: user

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./user-srv
```

Build a docker image
```
make docker
```