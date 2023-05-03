package test

import (
	"fmt"
	"testing"
	"tool1/cmd/wallets"
)

func TestNewClient(t *testing.T) {
	client := wallets.NewClient(fujiChainRPC)
	client, err := client.NewWallet()
	if err != nil {
		t.Error(err)
	}

	addr, err := client.Wallet.Address(idWalletTest1)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("rpc: :", client.RPC)
	fmt.Printf("Address id %d: %s\n", idWalletTest1, addr)
}

func TestNewClientFromMnemonic(t *testing.T) {
	client := wallets.NewClient(fujiChainRPC)
	client, err := client.NewWalletFromMnemonic(mnemonicTest)
	if err != nil {
		t.Error(err)
	}

	addr, err := client.Wallet.Address(idWalletTest1)
	if err != nil {
		t.Error(err)
	}

	fmt.Println("rpc: :", client.RPC)
	fmt.Printf("Address id %d: %s\n", idWalletTest1, addr)
}

func TestTransferNative(t *testing.T) {
	client := wallets.NewClient(fujiChainRPC)
	client, err := client.NewWalletFromMnemonic(mnemonicTest)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(client.Wallet.Address(1))

	err = client.TransferNativeToken(idWalletTest1, addressTest, Amount())
	if err != nil {
		t.Error(err)
	}
}

func TestTransferERC20Token(t *testing.T) {
	client := wallets.NewClient(fujiChainRPC)
	client, err := client.NewWalletFromMnemonic(mnemonicTest)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(client.Wallet.Address(1))

	err = client.TransferERC20Token(idWalletTest1, addressTest, TUSDContract, Amount1())
	if err != nil {
		t.Error(err)
	}
}
