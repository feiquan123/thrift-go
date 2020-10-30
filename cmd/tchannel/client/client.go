package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"time"

	tchannel "github.com/uber/tchannel-go"
	"github.com/uber/tchannel-go/thrift"

	"localhost.com/thrift-go/gen/tchannel/echo"
)

func main() {
	addr := "localhost:7878"

	// need creaet server name
	if err := runEchoClient("EchoServer", addr); err != nil {
		log.Fatalf("runEchoClient failed: %v", err)
	}

	go listenConsole()

	// Run for 10 seconds, then stop
	time.Sleep(time.Second * 10)
}

// set client
func runEchoClient(echoService string, addr string) error {
	tchan, err := tchannel.NewChannel("EchoClient", optsFor("EchoClient"))
	if err != nil {
		return err
	}
	tchan.Peers().Add(addr)
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

// opts
func optsFor(processName string) *tchannel.ChannelOptions {
	return &tchannel.ChannelOptions{
		ProcessName: processName,
		Logger:      tchannel.NewLevelLogger(tchannel.SimpleLogger, tchannel.LogLevelWarn),
	}
}
