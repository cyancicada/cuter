package demo

import (
	// "github.com/yakaa/cuterapplication/test/demo/remotedemo"
	"github.com/yakaa/cuter/lib/rpcx"
	"errors"
)

type (
	DemoRpcClient struct {
		cli *rpcx.RpcClient
	}
	DemoRpcResponse struct {
		Token string `json:"token"`
	}
)

var ErrNoRegionClient = errors.New("not found client")

func NewDemoRpcClient(cli *rpcx.RpcClient) *DemoRpcClient {
	return &DemoRpcClient{
		cli: cli,
	}
}

func (m *DemoRpcClient) DemoFunc(t string) (*DemoRpcResponse, error) {
	_, exist := m.cli.Next()
	if !exist {
		return nil, ErrNoRegionClient
	}
	return &DemoRpcResponse{Token: ""}, nil
}
