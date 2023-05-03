package wallets

import (
	"math/big"
	"tool1/app/contract"
)

func (c *Client) TransferNativeToken(id int, to string, amount *big.Int) error {
	wl, err := c.Wallet.Wallet(id)
	if err != nil {
		return err
	}

	transferCli := contract.NewClient(c.RPC)

	err = transferCli.NativeToken(wl.PrivateKey, to, amount)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) TransferERC20Token(id int, to, ct string, amount *big.Int) error {
	wl, err := c.Wallet.Wallet(id)
	if err != nil {
		return err
	}

	transferCli := contract.NewClient(c.RPC)

	err = transferCli.ERC20Token(wl.PrivateKey, to, ct, amount)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) MintTusd(id int) error {
	wl, err := c.Wallet.Wallet(id)
	if err != nil {
		return err
	}

	ctClient := contract.NewClient(c.RPC)

	err = ctClient.MintTUSD(wl.PrivateKey, wl.PublicKey)
	if err != nil {
		return err
	}

	return nil
}
