package test

import (
	"log"
	"math/big"
)

const (
	mnemonicTest  = "coach rapid rail fat connect pigeon plate museum salt practice shy double"
	fujiChainRPC  = "https://api.avax-test.network/ext/bc/C/rpc"
	idWalletTest1 = 1
	amount        = "10000000000000000" // 0.01 token
	amount1       = "10000"             // 0.01 token
	addressTest   = "0xEd6A2Dd9A563A3a682638FA543F63fB59bc9800a"
	TUSDContract  = "0xd00b9bbc6edc3953ec502d73e7fa7c59f628d947"
)

func Amount() *big.Int {
	value, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		log.Println("SetString: error")
		return nil
	}
	return value
}

func Amount1() *big.Int {
	value, ok := new(big.Int).SetString(amount1, 10)
	if !ok {
		log.Println("SetString: error")
		return nil
	}
	return value
}
