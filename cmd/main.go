package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/params"
	"github.com/joho/godotenv"
	ens "github.com/wealdtech/go-ens/v3"
)

const domainName = "example.eth"

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// Connect to the Ethereum network
	ec, err := ethclient.Dial(os.Getenv("INFURA_URL"))
	if err != nil {
		panic(err)
	}

	// Resolve ENS name into an Etheruem address
	ethAddr, err := ens.Resolve(ec, domainName)
	if err != nil {
		panic(err)
	}

	// Print the resolved Ethereum address
	fmt.Println("Resolved Ethereum address:", ethAddr.Hex())

	// Get the latest block number
	blockNumber, err := ec.BlockNumber(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Latest block number:", blockNumber)

	// Get the current balance of the Ethereum address using the same block number
	balanceWei, err := ec.BalanceAt(ctx, ethAddr, big.NewInt(int64(blockNumber)))
	if err != nil {
		panic(err)
	}
	fmt.Println("Current balance in WEI:", balanceWei.String())
	fmt.Println("Current balance in ETH:", weiToEther(balanceWei))

	// Get the transaction count of the Ethereum address
	nonce, err := ec.PendingNonceAt(ctx, ethAddr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction count:", nonce)
}

func weiToEther(wei *big.Int) *big.Float {
	f := new(big.Float)
	f.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	f.SetMode(big.ToNearestEven)
	fWei := new(big.Float)
	fWei.SetPrec(236) //  IEEE 754 octuple-precision binary floating-point format: binary256
	fWei.SetMode(big.ToNearestEven)
	return f.Quo(fWei.SetInt(wei), big.NewFloat(params.Ether))
}
