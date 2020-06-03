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

func (c *Client) GetQRIS(data *GetQRISParam) (*QRIS, *xendit.Error) {
	return c.GetQRISWithContext(context.Background(), data)
}

func (c *Client) GetQRISWithContext(ctx context.Context, data *GetQRISParam) (*QRIS, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &QRIS{}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/qr_codes/%s", c.Opt.XenditURL, data.ExternalID),
		c.Opt.SecretKey,
		nil,
		nil,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (c *Client) SimulatePayment(data *PaymentParam) (*PaymentResponse, *xendit.Error) {
	return c.SimulatePaymentWithContext(context.Background(), data)
}

func (c *Client) SimulatePaymentWithContext(ctx context.Context, data *PaymentParam) (*PaymentResponse, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &PaymentResponse{}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/qr_codes/%s/payments/simulate", c.Opt.XenditURL, data.ExternalID),
		c.Opt.SecretKey,
		nil,
		data,
		response,
	)

	if err != nil {
		return nil, err
	}

	return response, nil
}
