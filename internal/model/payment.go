package model

import "time"

type Payment struct {
	Id           int
	RequestId    string    `json:"request_id"`
	Currency     string    `json:"currency"`
	Transaction  string    `json:"transaction"`
	Provider     string    `json:"provider"`
	Amount       int       `json:"amount"`
	PaymentDt    time.Time `json:"payment_dt"`
	Bank         string    `json:"bank"`
	DeliveryCost int       `json:"delivery_cost"`
	GoodsTotal   int       `json:"goods_total"`
	CustomFee    int       `json:"custom_fee"`
}
