package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"runtime"
	"strings"
	"time"

	tchannel "github.com/uber/tchannel-go"
	"github.com/uber/tchannel-go/thrift"

	"localhost.com/thrift-go/gen/tchannel/echo"
)

func main() {
	var (
		listener net.Listener
		err      error
	)

	if listener, err = setupServer(); err != nil {
		log.Fatalf("setupServer failed: %v", err)
	}

	// need creaet server name
	if err := runEchoClient("EchoServer", listener.Addr()); err != nil {
		log.Fatalf("runEchoClient failed: %v", err)
	}

	go listenConsole()

	// Run for 10 seconds, then stop
	time.Sleep(time.Second * 10)
}

// set server
func setupServer() (net.Listener, error) {
	tchan, err := tchannel.NewChannel("EchoServer", optsFor("EchoServer"))
	if err != nil {
		return nil, err
	}

	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, err
	}

	server := thrift.NewServer(tchan)
	server.Register(echo.NewTChanEchoServer(&EchoHandler{}))

	// Serve will set the local peer info, and start accepting sockets in a separate goroutine.
	tchan.Serve(listener)
	return listener, nil
}

// set client
func runEchoClient(echoService string, addr net.Addr) error {
	tchan, err := tchannel.NewChannel("EchoClient", optsFor("EchoClient"))
	if err != nil {
		return err
	}
	tchan.Peers().Add(addr.String())
	tclient := thrift.NewClient(tchan, echoService, nil)
	client := echo.NewTChanEchoClient(tclient)

	go func() {
		for {
			ctx, cancel := thrift.NewContext(time.Second)
			req := echo.NewEchoReq()
			req.Msg = "Hello"
			res, err := client.Echo(ctx, req)
			log.Println("Echo(Hello) = ", res, ", err: ", err)
			cancel()
			time.Sleep(time.Second)
		}
	}()

	return nil
}

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

// EchoHandler
type EchoHandler struct{}

func (e *EchoHandler) Echo(ctx thrift.Context, req *echo.EchoReq) (*echo.EchoRes, error) {
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
