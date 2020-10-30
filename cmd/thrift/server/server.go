package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"

	"localhost.com/thrift-go/gen/thrift/echo"
)

type echoHandler struct{}

func (e *echoHandler) Echo(ctx context.Context, req *echo.EchoReq) (r *echo.EchoRes, err error) {
	fmt.Printf("message from client: %s\n", req.GetMsg())

	res := &echo.EchoRes{
		Msg: "sucess",
	}

	return res, nil
}

func main() {
	transport, err := thrift.NewTServerSocket(":9898")
	if err != nil {
		panic(err)
	}

	handler := &echoHandler{}
	processor := echo.NewEchoProcessor(handler)

	// 通信协议
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	if err := server.Serve(); err != nil {
		panic(err)
	}
}
