package qris

import (
	"context"
	"github.com/xendit/xendit-go"
)

type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

func CreateQRIS(data *CreateQRISParams) (*QRIS, *xendit.Error) {
	return CreateQRISWithContext(context.Background(), data)
}

func CreateQRISWithContext(ctx context.Context, data *CreateQRISParams) (*QRIS, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateQRISWithContext(ctx, data)

}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
