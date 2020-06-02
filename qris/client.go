package qris

import (
	"context"
	"fmt"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/utils/validator"
	"net/http"
)

func (c *Client) CreateQRIS(data *CreateQRISParams) (*QRIS, *xendit.Error) {
	return c.CreateQRISWithContext(context.Background(), data)
}

func (c *Client) CreateQRISWithContext(ctx context.Context, data *CreateQRISParams) (*QRIS, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &QRIS{}
	header := &http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/qr_codes", c.Opt.XenditURL),
		c.Opt.SecretKey,
		header,
		data,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
