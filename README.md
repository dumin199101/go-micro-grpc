
#### go-micro v2 默认是 GRPC 通信

#### 确保开启了 go module 模式
```
# 代理
GOPROXY=https://goproxy.io,direct
# 当前项目下
go mod tidy
```

#### 安装 micro v2 版本 v2.9.3  [可选]
```
https://github.com/micro/micro/releases/tag/v2.9.3
```
#### 调用 micro 生成代码 [可选]
> 进入当前项目目录
```
# 默认情况下 micro 生成的代码会放到 GOPATH/src 中，通过配置 --gopath=false 可以选择在当前目录下
micro new [--gopath=false] [--namespace=xxx] [--type=xxx] xxx

# micro 参数说明
--gopath    是否在 $GOPATH 目录下创建
--namespace 默认 go.micro
--type=     默认 service
xxx         项目相对路径
```

#### 安装 go-micro v2 版本 v2.9.1
```
go get github.com/micro/go-micro/v2
```

#### 生成工具 - protoc 插件
```
# protoc v3.11.4
https://github.com/protocolbuffers/protobuf/releases

# protoc-gen-go v1.4.3
go get github.com/golang/protobuf/protoc-gen-go@v1.4.3

# protoc-gen-micro v2.9.3
go get github.com/micro/micro/v2/cmd/protoc-gen-micro
```

#### proto 文件编译
```
protoc --micro_out=. --go_out=. *.proto

protoc 参数说明
--proto_path  proto 文件目录(简写 -I 如果没有指定参数，则在当前目录进行搜索)
--micro_out   生成的 micro 源码保存目录
--go_out      生成的 go 源码保存目录
```

#### etcd 单机版 v3.4.15
```
./etcd --listen-client-urls=http://0.0.0.0:2379 --advertise-client-urls=http://0.0.0.0:2379
```

#### 启动服务端
> 如果用 etcd 不需要在代码里写注册 etcd 的信息，只需要命令行添加即可

```
./go-micro-grpc --registry=etcd --registry_address=127.0.0.1:2379 --server_address=127.0.0.1:5501
```

#### 启动客户端
> 如果用 etcd 不需要在代码里写注册 etcd 的信息，只需要命令行添加即可
```
./client --registry=etcd --registry_address=127.0.0.1:2379    # 启用 etcd 服务发现和注册
```

#### 服务管理状态 micro web  [可选]
> http://localhost:8082
```
micro --registry=etcd --registry_address=127.0.0.1:2379 web
```

#### go-plugins consul 插件下载 [可选]
> 注意：因为 go-micro v2 后，去除了对 consul 的支持，使用插件
```
// consul 插件下载
go get github.com/micro/go-plugins/registry/consul/v2

// go 导入
import _ "github.com/micro/go-plugins/registry/consul/v2"
```

#### go-plugins consul 启动服务端与客户端 [可选]
```
--registry=consul --registry_address=127.0.0.1:8500    # 启用 consul 服务发现和注册
--server_address=127.0.0.1:5501                        # 给服务端添加端口绑定
```

#### go-plugins hystrix 断路器插件下载 [可选]
> 在客户端写入代码
```
go get github.com/micro/go-plugins/wrapper/breaker/hystrix/v2
```
