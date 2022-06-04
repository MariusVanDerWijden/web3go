package main

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	SK   = "0xaf5ead4413ff4b78bc94191a2926ae9ccbec86ce099d65aaf469e9eb1a0fa87f"
	ADDR = "0x6177843db3138ae69679A54b95cf345ED759450d"
)

func main() {
	// Use a local node:
	cl, err := ethclient.Dial("/tmp/geth.ipc")
	if err != nil {
		panic(err)
	}
	_ = cl
	// Prefund your address
	// eth.sendTransaction({from:personal.listAccounts[0], to:"0x6177843db3138ae69679A54b95cf345ED759450d", value: "10000000000000000"})
	// Send a transaction
	// sendTransaction(cl)
}

// createKeyPair creates a secret, public key pair, and corresponding address.
func createKeyPair() (*ecdsa.PrivateKey, *ecdsa.PublicKey, common.Address) {
	sk, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}
	addr := crypto.PubkeyToAddress(sk.PublicKey)
	return sk, &sk.PublicKey, addr
}
