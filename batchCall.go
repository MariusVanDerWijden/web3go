package main

import (
	"context"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/rpc"
)

func batchGetTransaction(url string) {
	client, err := rpc.Dial(url)
	if err != nil {
		panic(err)
	}
	var (
		results = make([]hexutil.Big, 2)
		reqs    = make([]rpc.BatchElem, 2)
	)
	// Set up the requests
	for i := range results {
		reqs[i] = rpc.BatchElem{
			Method: "eth_chainId",
			Args:   nil,
			Result: &results[i],
		}
	}
	// Send the batch to the client
	err = client.BatchCallContext(context.Background(), reqs)
	if err != nil {
		panic(err)
	}
	for i, res := range results {
		fmt.Printf("Request: %v Result: %v\n", i, res.ToInt())
	}
}
