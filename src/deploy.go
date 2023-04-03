package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	// "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xiaofeng-tech/contract-demo/contract"
)

func main() {
	// Connect to an Ethereum client
    client, err := ethclient.Dial("https://sepolia.infura.io/v3/9b3dc42eba134e8f99a65646f3923cb2")
    if err != nil {
        log.Fatal(err)
    }

    // Get the private key for the account that will deploy the contract
    privateKey, err := crypto.HexToECDSA("xxx")
    if err != nil {
        log.Fatal(err)
    }

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
        log.Fatal("error casting public key to ECDSA")
    }

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        log.Fatal(err)
    }

	auth := bind.NewKeyedTransactor(privateKey)
    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)     // in wei
    auth.GasLimit = uint64(1000000) // in units
    auth.GasPrice = gasPrice

	// input := "1.0"
	address, tx, instance, err := contract.DeployContract(auth, client)
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
    fmt.Println(tx.Hash().Hex())

	_ = instance
}