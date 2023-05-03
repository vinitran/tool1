package wallets

import "tool1/app/hdwallet"

type Client struct {
	Wallet *hdwallet.HdWallet `json:"wallet"`
	RPC    string             `json:"rpc"`
}

func NewClient(rpc string) *Client {
	return &Client{
		Wallet: nil,
		RPC:    rpc,
	}
}

func (c *Client) NewWallet() (*Client, error) {
	hdwl, err := hdwallet.NewHdWallet()
	if err != nil {
		return nil, err
	}

	return &Client{
		Wallet: hdwl,
		RPC:    c.RPC,
	}, nil
}

func (c *Client) NewWalletFromMnemonic(mnemonic string) (*Client, error) {
	hdwl, err := hdwallet.NewHdWalletByMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}

	return &Client{
		Wallet: hdwl,
		RPC:    c.RPC,
	}, nil
}
