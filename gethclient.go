package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/ethclient/gethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

func newGethClient(url string) *gethclient.Client {
	client, err := rpc.Dial(url)
	if err != nil {
		panic(err)
	}
	return gethclient.New(client)
}

func queryGethClient(url string) error {
	client := newGethClient(url)
	mem, err := client.MemStats(context.Background())
	if err != nil {
		return err
	}
	fmt.Printf("Memory consumption: %v\n", mem.Alloc)
	return nil
}
