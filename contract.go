package main

import (
	"context"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/mariusvanderwijden/web3go/contracts"
)

func DeployContract(cl *ethclient.Client) error {
	var (
		sk = crypto.ToECDSAUnsafe(common.FromHex(SK))
	)
	// Retrieve the chainid (needed for signer)
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		return err
	}

	// Create the transactOpts (signer)
	transactOpts, err := bind.NewKeyedTransactorWithChainID(sk, chainid)
	if err != nil {
		return err
	}

	// Deploy a CoolContract
	addr, tx, contract, err := contracts.DeployCoolContract(transactOpts, cl)
	if err != nil {
		return err
	}
	_ = addr
	_ = contract

	// Wait until the contract is deployed
	addr, err = bind.WaitDeployed(context.Background(), cl, tx)
	if err != nil {
		return err
	}
	fmt.Printf("CoolContract deployed at %v\n", addr)

	// Query the balance of the contract
	if err := callSeeBalance(contract); err != nil {
		return err
	}

	// Call the deposit function
	if err := callDeposit(cl, contract, transactOpts); err != nil {
		return err
	}

	// Query the balance of the contract again
	if err := callSeeBalance(contract); err != nil {
		return err
	}

	// filter deposit events
	if err := filterDepositEvents(contract); err != nil {
		return err
	}

	return nil
}

func callDeposit(cl *ethclient.Client, contract *contracts.CoolContract, txOpts *bind.TransactOpts) error {
	txOpts.Value = new(big.Int).Mul(big.NewInt(1), big.NewInt(params.GWei))
	// Interact with the contract
	tx, err := contract.Deposit(txOpts)
	if err != nil {
		return err
	}
	// Wait until the transaction is mined
	receipt, err := bind.WaitMined(context.Background(), cl, tx)
	if err != nil {
		return err
	}
	_ = receipt
	return nil
}

func callSeeBalance(contract *contracts.CoolContract) error {
	callOpts := &bind.CallOpts{Context: context.Background(), Pending: false}
	bal, err := contract.SeeBalance(callOpts)
	if err != nil {
		return err
	}
	fmt.Printf("CoolContract balance: %v\n", bal)
	return nil
}

func filterDepositEvents(contract *contracts.CoolContract) error {
	// Filter for a Deposited event
	filterOpts := &bind.FilterOpts{Context: context.Background(), Start: 0, End: nil}
	itr, err := contract.FilterDeposited(filterOpts)
	if err != nil {
		return err
	}
	// Loop over all found events
	for itr.Next() {
		event := itr.Event
		// Print out all caller addresses
		// TODO: add a value field to the event
		fmt.Printf("Address %v deposited %v wei\n", event.Addr.Hex(), 0)
	}
	return nil
}

func watchDepositEvents(contract *contracts.CoolContract) error {
	// Watch for a Deposited event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}
	// Setup a channel for results
	channel := make(chan *contracts.CoolContractDeposited)
	// Start a goroutine which watches new events
	go func() {
		sub, err := contract.WatchDeposited(watchOpts, channel)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()
	}()
	// Receive events from the channel
	event := <-channel
	fmt.Printf("Event received: %v\n", event.Addr)
	return nil
}
