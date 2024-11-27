package bknd

import (
	"github.com/go-resty/resty/v2"
)

type RestyClient struct {
	client *resty.Client
}

func NewRestyClient() *RestyClient {
	return &RestyClient{
		client: resty.New(),
	}
}
