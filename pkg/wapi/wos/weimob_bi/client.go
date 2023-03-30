package weimob_bi

import (
	"github.com/weimob-tech/go-ability-sdk/pkg/sdk/api"
)

type Client struct {
	api.Client
}

func NewClientWithClientKey(clientId, clientSecret string) (client *Client, err error) {
	client = &Client{}
	err = client.InitWithAccessKey(clientId, clientSecret)
	return
}
