package client

import "github.com/ybbus/jsonrpc/v3"

var rpcClient jsonrpc.RPCClient

func InitBinder(endpoint string) {
	rpcClient = jsonrpc.NewClient(endpoint)
}

func BinderClient() *jsonrpc.RPCClient {
	return &rpcClient
}
