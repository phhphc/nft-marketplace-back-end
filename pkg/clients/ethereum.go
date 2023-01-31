package clients

import "github.com/ethereum/go-ethereum/ethclient"

type EthClient struct {
	*ethclient.Client
}

func NewEthClient(ethUrl string) (*EthClient, error) {
	client, err := ethclient.Dial(ethUrl)
	if err != nil {
		return nil, err
	}
	return &EthClient{
		Client: client,
	}, nil
}

func (c *EthClient) Disconnect() {
	c.Client.Close()
}
