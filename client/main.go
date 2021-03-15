package main

import (
	"context"
	"fmt"

	hystrixgo "github.com/afex/hystrix-go/hystrix"
	"github.com/laixhe/go-micro-grpc/protorpc"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
)

const (
	// 当前服务-服务名
	ServerName = "test.user.client"
	// 用户服务-服务名
	UserServerName = "test.user.server"
)

func main() {

	// 其实这里不需要在代码里写注册 etcd 的信息，只需要命令行添加即可 --registry=etcd --registry_address=127.0.0.1:2379
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:2379",
		}
	})

	// 创建服务
	service := micro.NewService(
		micro.Name(ServerName),
		micro.Version("v1.0.0"),
		micro.Registry(reg),
		micro.WrapClient(
			hystrix.NewClientWrapper(),
		),
	)

	// 自定义全局默认超时时间和最大并发数
	hystrixgo.DefaultTimeout = 2000
	hystrixgo.DefaultMaxConcurrent = 20

	// 初始化服务
	service.Init()

	// 获取用户 服务
	UserRPC := protorpc.NewUserService(UserServerName, service.Client())

	res, err := UserRPC.GetUser(context.Background(), &protorpc.GetUserRequest{
		Userid: 123654,
	})

	fmt.Println(res, err)
}
