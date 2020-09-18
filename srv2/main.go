package main

import (
	"context"
	"fmt"
	pb "micro-hello/proto/hello"
	"time"

	"github.com/micro/go-micro/v2"
)

func main() {
	srv := micro.NewService(

		micro.Address("127.0.0.1:6789"),
		micro.Name("hello.srv2"),
	)

	client := pb.NewHelloService("hello.srv1", srv.Client())

	for i := 0; i < 10; i++ {
		resp, err := client.Greeter(context.TODO(), &pb.Request{Msg: "hello"})
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(resp.Msg)
		time.Sleep(10)
	}
}
