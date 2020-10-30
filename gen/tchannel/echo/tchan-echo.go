// @generated Code generated by thrift-gen. Do not modify.

// Package echo is generated code used to make or handle TChannel calls using Thrift.
package echo

import (
	"fmt"

	athrift "github.com/apache/thrift/lib/go/thrift"
	"github.com/uber/tchannel-go/thrift"
)

// Interfaces for the service and client for the services defined in the IDL.

// TChanEcho is the interface that defines the server handler and client interface.
type TChanEcho interface {
	Echo(ctx thrift.Context, req *EchoReq) (*EchoRes, error)
}

// Implementation of a client and service handler.

type tchanEchoClient struct {
	thriftService string
	client        thrift.TChanClient
}

func NewTChanEchoInheritedClient(thriftService string, client thrift.TChanClient) *tchanEchoClient {
	return &tchanEchoClient{
		thriftService,
		client,
	}
}

// NewTChanEchoClient creates a client that can be used to make remote calls.
func NewTChanEchoClient(client thrift.TChanClient) TChanEcho {
	return NewTChanEchoInheritedClient("Echo", client)
}

func (c *tchanEchoClient) Echo(ctx thrift.Context, req *EchoReq) (*EchoRes, error) {
	var resp EchoEchoResult
	args := EchoEchoArgs{
		Req: req,
	}
	success, err := c.client.Call(ctx, c.thriftService, "echo", &args, &resp)
	if err == nil && !success {
		switch {
		default:
			err = fmt.Errorf("received no result or unknown exception for echo")
		}
	}

	return resp.GetSuccess(), err
}

type tchanEchoServer struct {
	handler TChanEcho
}

// NewTChanEchoServer wraps a handler for TChanEcho so it can be
// registered with a thrift.Server.
func NewTChanEchoServer(handler TChanEcho) thrift.TChanServer {
	return &tchanEchoServer{
		handler,
	}
}

func (s *tchanEchoServer) Service() string {
	return "Echo"
}

func (s *tchanEchoServer) Methods() []string {
	return []string{
		"echo",
	}
}

func (s *tchanEchoServer) Handle(ctx thrift.Context, methodName string, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	switch methodName {
	case "echo":
		return s.handleEcho(ctx, protocol)

	default:
		return false, nil, fmt.Errorf("method %v not found in service %v", methodName, s.Service())
	}
}

func (s *tchanEchoServer) handleEcho(ctx thrift.Context, protocol athrift.TProtocol) (bool, athrift.TStruct, error) {
	var req EchoEchoArgs
	var res EchoEchoResult

	if err := req.Read(protocol); err != nil {
		return false, nil, err
	}

	r, err :=
		s.handler.Echo(ctx, req.Req)

	if err != nil {
		return false, nil, err
	} else {
		res.Success = r
	}

	return err == nil, &res, nil
}