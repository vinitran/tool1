package contract

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"tool1/app/abi/erc20"
)

type Client struct {
	Client *ethclient.Client
}

func NewClient(rpc string) *Client {
	client, err := ethclient.Dial(rpc)
	if err != nil {
		fmt.Println("error: cant connect to client")
		log.Fatal(err)
	}
	return &Client{Client: client}
}

func (c *Client) NativeToken(privateK, to string, amount *big.Int) error {
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

	gasLimit := uint64(gasLimitNativeToken) // in units
	gasPrice, err := c.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return err
	}

	toAddress := common.HexToAddress(to)
	var data []byte
	tx := types.NewTx(&types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    amount,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	})

	chainID, err := c.Client.NetworkID(context.Background())
	if err != nil {
		return err
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		return err
	}

	err = c.Client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return err
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
	return nil
}

func (c *Client) ERC20Token(privateK, to, contract string, amount *big.Int) error {
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

	toAddress := common.HexToAddress(to)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)                 // in wei
	auth.GasLimit = uint64(gasLimitURC20Token) // in units
	auth.GasPrice = gasPrice

	tokenContract := common.HexToAddress(contract)
	instance, err := erc20.NewToken(tokenContract, c.Client)
	if err != nil {
		return err
	}

	tx, err := instance.Transfer(auth, toAddress, amount)
	if err != nil {
		return err
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex())
	return nil
}
