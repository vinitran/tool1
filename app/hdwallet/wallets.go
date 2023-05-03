package hdwallet

import (
	"encoding/json"
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
)

const bits = 128

type HdWallet struct {
	HdWallet *hdwallet.Wallet `json:"hd_wallet"`
}

type Wallet struct {
	PrivateKey string `json:"private_key"`
	PublicKey  string `json:"public_key"`
	Address    string `json:"address"`
}

func (h *HdWallet) String() string {
	wallet, err := json.Marshal(h)
	if err != nil {
		log.Println("Can not marshall wallets .")
		return ""
	}

	return string(wallet)
}

func (wl *Wallet) String() string {
	wallet, err := json.Marshal(wl)
	if err != nil {
		log.Println("Can not marshall wallets .")
		return ""
	}

	return string(wallet)
}

func NewHdWallet() (*HdWallet, error) {
	entropy, err := hdwallet.NewEntropy(bits)
	if err != nil {
		return nil, err
	}

	mnemonic, err := hdwallet.NewMnemonicFromEntropy(entropy)
	if err != nil {
		return nil, err
	}

	log.Printf("Created new seed: %s\n", mnemonic)

	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	hdWallet := HdWallet{
		HdWallet: wallet,
	}

	return &hdWallet, nil
}

func NewHdWalletByMnemonic(mnemonic string) (*HdWallet, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	log.Printf("Load data from mnemonic successfully!")

	hdWallet := HdWallet{
		HdWallet: wallet,
	}

	return &hdWallet, nil
}

func (h *HdWallet) WalletByDerivationPath(derivationPath string) (*Wallet, error) {
	path := hdwallet.MustParseDerivationPath(derivationPath)
	account, err := h.HdWallet.Derive(path, false)
	if err != nil {
		return nil, err
	}

	privateKey, err := h.HdWallet.PrivateKeyHex(account)
	if err != nil {
		return nil, err
	}

	publicKey, err := h.HdWallet.PublicKeyHex(account)
	if err != nil {
		return nil, err
	}

	address, err := h.HdWallet.AddressHex(account)
	if err != nil {
		return nil, err
	}

	wl := Wallet{
		PrivateKey: privateKey,
		PublicKey:  publicKey,
		Address:    address,
	}
	return &wl, nil
}

func (h *HdWallet) Wallet(id int) (*Wallet, error) {
	derivationPath := fmt.Sprintf("m/44'/0'/%d'/0/%d", 60, id)
	wl, err := h.WalletByDerivationPath(derivationPath)
	if err != nil {
		return nil, err
	}
	return wl, nil
}

func (h *HdWallet) PrivateKey(id int) (string, error) {
	wallet, err := h.Wallet(id)
	if err != nil {
		return "", err
	}

	return wallet.PrivateKey, nil
}

func (h *HdWallet) Address(id int) (string, error) {
	wallet, err := h.Wallet(id)
	if err != nil {
		return "", err
	}
	return wallet.Address, nil
}
