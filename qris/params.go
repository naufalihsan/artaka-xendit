package qris

type QRIS struct {
	ID          string       `json:"id"`
	ExternalID  string       `json:"external_id"`
	Amount      float64      `json:"amount"`
	QRString    string       `json:"qr_string"`
	CallbackURL string       `json:"callback_url"`
	Type        QRISTypeEnum `json:"type"`
	Status      string       `json:"status"`
	Created     string       `json:"created"`
	Updated     string       `json:"updated"`
}

type PaymentResponse struct {
	ID      string  `json:"id"`
	Amount  float64 `json:"amount"`
	Created string  `json:"created"`
	QRCode  QRIS    `json:"qr_code"`
	Status  string  `json:"status"`
}

type CreateQRISParams struct {
	ForUserID   string       `json:"-"`
	ExternalID  string       `json:"external_id" validate:"required"`
	Type        QRISTypeEnum `json:"type" validate:"required"`
	CallbackURL string       `json:"callback_url",validate:"required"`
	Amount      float64      `json:"amount,omitempty"`
}

type GetQRISParam struct {
	ExternalID string `json:"external_id" validate:"required"`
}

type PaymentParam struct {
	ExternalID string  `json:"external_id" validate:"required"`
	Amount     float64 `json:"amount,omitempty"`
}

type QRISTypeEnum string

const (
	QRISTypeDYNAMIC QRISTypeEnum = "DYNAMIC"
	QRISTypeSTATIC  QRISTypeEnum = "STATIC"
)
