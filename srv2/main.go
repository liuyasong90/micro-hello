package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-plugins/registry/consul/v2"
	pb "micro-hello/proto/hello"
	"time"
)

func main() {
	srv := micro.NewService(

		micro.Name("hello.srv2"),
		micro.Registry(consul.NewRegistry(
			registry.Addrs("101.200.129.72:8500"),
			)),
		micro.RegisterTTL(time.Second*30),
		micro.RegisterInterval(time.Second*10),
	)

	client := pb.NewHelloService("hello.srv1.01", srv.Client())

	for i := 0; i < 10; i++ {
		resp, err := client.Greeter(context.TODO(), &pb.Request{Msg: "hello"})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Msg)
		time.Sleep(10)
	}
}
