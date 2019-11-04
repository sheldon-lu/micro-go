# micros
基于go-micro的用户接口微服务demo


```shell 
git clone https://github.com/sheldon-lu/micros.git
## 默认使用注册方式为mdns，windows无法使用，所以改成consul
## mac或者linux可直接使用mdns，也可以使用consul
## 安装请自行下载consul，启动consul
consul agent -dev

## 启动微服务
cd micros/user
## 因为是微服务，每个服务都是独立的模块，需要单独起服务。
go run main.go -registry=consul
## 如果有api服务代码先起api服务，再micro api，
## 或者直接起api接口,指定namespace
micro -registry=consul api --namespace=go.micro.srv
## --handler=api ---> handler参数api/web/rpc..
```
