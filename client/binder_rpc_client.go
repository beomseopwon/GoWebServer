package client

import "github.com/ybbus/jsonrpc/v3"

var binderClient jsonrpc.RPCClient

func InitBinder(endpoint string) {
	binderClient = jsonrpc.NewClient(endpoint)
}

func BinderClient() *jsonrpc.RPCClient {
	return &binderClient
}
