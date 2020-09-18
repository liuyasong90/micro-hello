package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	pb "micro-hello/proto/hello"
)

type Handler struct {
}

func (h Handler) Greeter(ctx context.Context, request *pb.Request, response *pb.Response) error {
	fmt.Println(request.Msg)
	response.Msg = request.Msg+" world"
	return nil
}

var _ pb.HelloHandler = (*Handler)(nil)

func main() {

	//声明服务
	srv := micro.NewService(

		micro.Registry(consul.NewRegistry(
			registry.Addrs("101.200.129.72:8500"),
		)),
		micro.Name("hello.srv1.01"),
		micro.Version("latest"),
	)

	//初始化
	srv.Init()

	//注册Handler
	h := Handler{}
	_ = pb.RegisterHelloHandler(srv.Server(), &h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
