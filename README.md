# micro-go
### 基于go-micro的demo项目，srv -> web -> apigateway模式
~~~ 
srv层 基于底层逻辑                   // go-micro
web层 基于面向外部接口与底层交互      // web层接口 使用了gin框架实现
api层 外部网关层                     // micro api --handler=web

micro call命令直接调用srv服务

目录结构：通用方法及配置目录
.
├── common-config            * 配置文件目录
│   └── config               * 配置类
│   │   └── config.go        * 初始化配置类
│   │   └── etcd.go          * etcd配置结构体 /配置中心
│   │   └── mysql.go         * mysql配置结构体
│   │   └── profiles.go      * 配置文件树辅助类
│   └── db                   * 数据库相关
│   │    └── db.go           * 初始化数据库
│   │    └── mysql.go        * mysql数据库相关
│   └── basic                * 初始化基础组件
├── common    
~~~


```shell
-- v1.0.0
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
cd micro-go/user

## 因为是微服务，每个服务都是独立的模块，需要单独起服务。
go run main.go -registry=consul


## 推荐三层架构。
## 如果有api服务代码启动api服务和micro api统一入口，
## micro api：(localhost:8080)-http访问入口,此作为api gateway网关
## api service：(go.micro.api.xxx)-对外暴露的API服务
## backend service：(go.micro.srv.xxx)-内网的后台服务
## 对应启动：
## micro -registry api   ## --handler=api /web/rpc/proxy/event
## --handler=api ---> handler参数api/web/rpc/proxy/event..
go run api/api.go
go run srv/main.go

## 或者直接起api接口,指定namespace,不指定会默认去go.micro.api中查找服务
micro -registry=consul api --namespace=go.micro.srv

```

补充：
API Handler
API处理器接收任何的HTTP请求，并且向前转发指定格式的RPC请求。

Content-Type: 支持任何类型
Body: 支持任何格式
Forward Format: 转发格式会按照go-micro中的go-api的方法，
[api.Request](https://github.com/asim/go-api/blob/master/proto/api.proto)\/api.Response

Path: 请求路径，/[service]/[method]
Resolver: 请求解析器，路径会被解析成服务与方法
Configure: 配置，在启动时指定--handler=api或在启动命令前指定环境变量MICRO_API_HANDLER=api

