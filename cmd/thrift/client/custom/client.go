package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/apache/thrift/lib/go/thrift"

	"localhost.com/thrift-go/gen/thrift/echo"
)

func main() {
	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "9898"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport, err := transportFactory.GetTransport(transport)
	client := echo.NewEchoClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "error opening socet to 127.0.0.1:9898 ", err)
		os.Exit(1)
	}
	defer transport.Close()

	req := &echo.EchoReq{Msg: "You are welcome."}
	ctx := context.Background()
	res, err := client.Echo(ctx, req)
	if err != nil {
		log.Println("Echo failed:", err)
		return
	}

	log.Println("response:", res.Msg)
	fmt.Println("well done!")
}
