package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

func main() {

	url := "https://godwoken-alphanet-v1.ckbapp.dev"
	//url:="https://api.bitstack.com/v1/wNFxbiJyQsSeLrX8RRCHi7NpRxrlErZk/DjShIqLishPCTB9HiMkPHXjUM9CNM9Na/ETH/mainnet"
	cli, err := ethclient.Dial(url)

	if err != nil {
		panic(err)
	}
	height,err  := cli.BlockNumber(context.Background())
	if err !=nil{
		return
	}
	for i := height; i > 0; i-- {
		blockMsg,err := cli.BlockByNumber(context.Background(), big.NewInt(int64(i)))
		if err != nil{
			fmt.Printf("blockNum%v,err%v\n",i,err)
			continue
		}
		fmt.Println(blockMsg)
	}

}