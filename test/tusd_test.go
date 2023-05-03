package test

import (
	"fmt"
	"testing"
	"tool1/cmd/wallets"
)

func TestMintTUSD(t *testing.T) {
	client := wallets.NewClient(fujiChainRPC)
	client, err := client.NewWalletFromMnemonic(mnemonicTest)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(client.Wallet.Address(1))

	err = client.MintTusd(idWalletTest1)
	if err != nil {
		t.Error(err)
	}
}
