package test

import (
	"fmt"
	"testing"
	"tool1/app/hdwallet"
)

func TestCreateHdWallet(t *testing.T) {
	_, err := hdwallet.NewHdWallet()
	if err != nil {
		t.Error(err)
	}
}

func TestGetHdWalletFromMnemonic(t *testing.T) {
	_, err := hdwallet.NewHdWalletByMnemonic(mnemonicTest)
	if err != nil {
		t.Error(err)
	}
}

func TestGetWalletFromId(t *testing.T) {
	hdwl, err := hdwallet.NewHdWalletByMnemonic(mnemonicTest)
	if err != nil {
		t.Error(err)
	}

	wl, err := hdwl.Wallet(0)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(wl.String())
}
