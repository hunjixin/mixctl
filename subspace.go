package main

import (
	"context"

	jsonrpc "github.com/ybbus/jsonrpc/v3"
)

type SubspaceClient struct {
	client jsonrpc.RPCClient
}

func NewSubpsaceClient(url string) *SubspaceClient {
	rpcClient := jsonrpc.NewClient(url)
	return &SubspaceClient{
		rpcClient,
	}
}

type FarmAppInfo struct {
	GenesisHash       string         `json:"genesisHash"`
	DsnBootstrapNodes []string       `json:"dsnBootstrapNodes"`
	Syncing           bool           `json:"syncing"`
	FarmTimeout       map[string]int `json:"farmingTimeout"`
}

func (c *SubspaceClient) GetFarmerAppInfo(ctx context.Context) (FarmAppInfo, error) {
	result := FarmAppInfo{}
	resp, err := c.client.Call(ctx, "subspace_getFarmerAppInfo")
	if err != nil {
		return result, err
	}

	err = resp.GetObject(&result)
	if err != nil {
		return result, err
	}
	return result, err
}
