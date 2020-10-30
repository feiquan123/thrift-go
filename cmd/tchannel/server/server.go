package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"

	tchannel "github.com/uber/tchannel-go"
	"github.com/uber/tchannel-go/thrift"

	"localhost.com/thrift-go/gen/tchannel/echo"
)

func main() {
	// no user listener
	var (
		listener net.Listener
		err      error
	)

	if listener, err = setupServer(); err != nil {
		log.Fatalf("setupServer failed: %v", err)
	}

	fmt.Println("Listener :", listener.Addr().String())

	go listenConsole()

	select {}
}

// set server
func setupServer() (net.Listener, error) {
	tchan, err := tchannel.NewChannel("EchoServer", optsFor("EchoServer"))
	if err != nil {
		return nil, err
	}

	listener, err := net.Listen("tcp", "localhost:7878")
	if err != nil {
		return nil, err
	}

	server := thrift.NewServer(tchan)
	server.Register(echo.NewTChanEchoServer(&echoHandler{}))

	// Serve will set the local peer info, and start accepting sockets in a separate goroutine.
	tchan.Serve(listener)
	return listener, nil
}

// listen consule input
func listenConsole() {
	rdr := bufio.NewReader(os.Stdin)
	for {
		line, _ := rdr.ReadString('\n')
		switch strings.TrimSpace(line) {
		case "s":
			printStack()
		default:
			fmt.Println("Unrecognized command:", line)
		}
	}
}

// printStack
func printStack() {
	buf := make([]byte, 10000)
	runtime.Stack(buf, true /* all */)
	fmt.Println("Stack:\n", string(buf))
}

// echoHandler
type echoHandler struct{}

func (e *echoHandler) Echo(ctx thrift.Context, req *echo.EchoReq) (*echo.EchoRes, error) {
	fmt.Printf("message from client: %s\n", req.GetMsg())

	res := &echo.EchoRes{
		Msg: "sucess",
	}

	return res, nil
}

// opts
func optsFor(processName string) *tchannel.ChannelOptions {
	return &tchannel.ChannelOptions{
		ProcessName: processName,
		Logger:      tchannel.NewLevelLogger(tchannel.SimpleLogger, tchannel.LogLevelWarn),
	}
}
