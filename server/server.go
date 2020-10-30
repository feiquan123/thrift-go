package main

import (
	"context"
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"

	"localhost.com/thrift-go/gen-go-13/echo"
)

type EchoServer struct{}

func (e *EchoServer) Echo(ctx context.Context, req *echo.EchoReq) (*echo.EchoRes, error) {
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

	handler := &EchoServer{}
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
