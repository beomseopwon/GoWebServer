package client

import "github.com/ybbus/jsonrpc/v3"

var vksClient jsonrpc.RPCClient

func InitVKS(endpoint string, jwt string) {
	vksClient = jsonrpc.NewClientWithOpts(endpoint, &jsonrpc.RPCClientOpts{
		CustomHeaders: map[string]string{
			"Authorization": "Bearere " + jwt,
		},
		AllowUnknownFields: true,
	})
}

func VKSClient() *jsonrpc.RPCClient {
	return &vksClient
}
