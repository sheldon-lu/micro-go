# User Service
# user服务模块结构
```sql
CREATE DATABASE `micros` /*!40100 DEFAULT CHARACTER SET utf8 */
```

```$xslt
# 生成接口代码（proto）：
protoc --proto_path=. --micro_out=. --go_out=. proto/user/user.proto


This is the User service

Generated with

```
处理底层逻辑 srv

micro new micro-go/user --namespace=lu.micro --alias=user --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: lu.micro.srv.user
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

```
补充：
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
├── common                   * 公共方法目录（数据库、redis等）

.
├── main.go                  * 主函数入口
├── plugin.go                * 自定义配置插件
├── handler                  * 处理器逻辑
│   └── user.go              * main.go srv与业务逻辑链接 ==> proto函数逻辑
├── model                    * 增加模型层，用于与数据库交换数据
│   └── user                 * 用户模型类
│   │   └── user.go          * 初始化用户模型类,封装获取用户数据类业务
│   └── user.go              * 初始化模型层
├── proto/user               * proto协议配置自定义
│   └── user.proto
├── Dockerfile
├── Makefile
└── README.md
```