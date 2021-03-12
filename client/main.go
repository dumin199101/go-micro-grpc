package main

import (
	"context"
	"fmt"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/registry/etcd"

	"github.com/laixhe/go-micro-grpc/protorpc"
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
		micro.Registry(reg),
		micro.Name("test.user.client"),
		micro.Version("v1.0.0"),
	)

	// 初始化服务
	service.Init()

	// 获取用户 服务
	UserRPC := protorpc.NewUserService("test.user.server", service.Client())

	res, err := UserRPC.GetUser(context.Background(), &protorpc.GetUserRequest{
		Userid: 123654,
	})

	fmt.Println(res, err)
}
