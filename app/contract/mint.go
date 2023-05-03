package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	TUSD "tool1/app/abi/tusd"
)

const (
	contractTUSD      = "0xd00B9BBC6EDC3953Ec502d73E7FA7C59f628d947"
	amountTUSD        = 15000000
	amountAVAXDefault = 1
)

func (c *Client) MintTUSD(privateK, to string) error {
	privateKey, err := crypto.HexToECDSA(privateK)
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return fmt.Errorf("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := c.Client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return err
	}

	gasPrice, err := c.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	chainID, err := c.Client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(amountAVAXDefault) // in wei
	auth.GasLimit = uint64(gasLimitURC20Token) // in units
	auth.GasPrice = gasPrice

	toAddress := common.HexToAddress(to)

	instance, err := TUSD.NewTUSD(common.HexToAddress(contractTUSD), c.Client)
	if err != nil {
		return err
	}

	tx, err := instance.Mint(auth, toAddress, big.NewInt(amountTUSD))
	if err != nil {
		return err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	return nil
}
