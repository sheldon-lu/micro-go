# micros
基于go-micro的用户接口微服务demo


```shell 
# 安装go-micro
# 使用如下指令安装
go get -u -v github.com/micro/micro
go get -u -v github.com/micro/go-micro

# 安装 protoc
# https://github.com/protocolbuffers/protobuf/releases下载
# 解压到目标文件架,我们以e:\dev为例
# 命令需要添加e:\dev\protoc-3.9.1-win64\bin到环境变量path

# 安装protoc-gen-micro插件
# 这个插件主要作用是通过.proto文件生成适用于go-micro的代码
go get -u -v github.com/micro/protoc-gen-micro
# 或者
go install github.com/micro/protoc-gen-micro
go install github.com/golang/protobuf/protoc-gen-go

# clone代码
git clone https://github.com/sheldon-lu/micros.git

## 默认使用注册方式为mdns，windows无法使用，所以改成consul
## mac或者linux可直接使用mdns，也可以使用consul
## 安装请自行下载consul，启动consul
consul agent -dev

## 启动微服务
cd micros/user

## 因为是微服务，每个服务都是独立的模块，需要单独起服务。
go run main.go -registry=consul

## 如果有api服务代码启动api服务和micro api统一入口，
## micro api：(localhost:8080)-http访问入口,此作为api gateway网关
## api service：(go.micro.api.xxx)-对外暴露的API服务
## backend service：(go.micro.srv.xxx)-内网的后台服务
## 对应启动：
micro -registry api --handler=api ## /web/rpc/proxy/event
## --handler=api ---> handler参数api/web/rpc/proxy/event..
go run api/api.go
go run srv/main.go

## 或者直接起api接口,指定namespace,不指定会默认去go.micro.api中查找服务
micro -registry=consul api --namespace=go.micro.srv

```
