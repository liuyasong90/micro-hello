package main

import (
	"context"
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	pb "micro-hello/proto/hello"

	"github.com/micro/go-micro/v2"
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
		micro.Name("hello.srv1"),
		micro.Version("latest"),
	)


	//初始化
	srv.Init(
		micro.BeforeStart(func() error {
			log.Log(1)
			return nil
		}),

		micro.AfterStart(func() error {
			log.Log(2)
			return nil
		}),
	)

	//注册Handler
	h := Handler{}
	_ = pb.RegisterHelloHandler(srv.Server(), &h)

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}
