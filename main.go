package main

import (
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/laixhe/go-micro-grpc/handler"
	"github.com/laixhe/go-micro-grpc/protorpc"
)

func main() {

	// 定义 etcd 的注册
	// 其实这里不需要在代码里写注册 etcd 的信息，只需要命令行添加即可 --registry=etcd --registry_address=127.0.0.1:2379
	reg := etcd.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:2379",
		}
	})

	// 创建服务
	service := micro.NewService(
		micro.Address("127.0.0.1:5501"), // 其实这里不需要在代码里写的，只需要命令行添加即可 --server_address=127.0.0.1:5501
		micro.Registry(reg),
		micro.Name("test.user.server"), // 当前微服务服务名
		micro.Version("v1.0.0"),
	)

	// 初始化服务
	service.Init()

	// 注册服务
	if err := protorpc.RegisterUserHandler(service.Server(), new(handler.UserService)); err != nil {
		fmt.Println("RegisterUserGetRPCHandler:", err)
		return
	}

	// 启动
	if err := service.Run(); err != nil {
		fmt.Println("Run:", err)
	}

}
