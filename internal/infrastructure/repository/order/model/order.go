package model

import (
	repoDelivery "tech-wb/internal/infrastructure/repository/delivery/model"
	repoItem "tech-wb/internal/infrastructure/repository/item/model"
	repoPayment "tech-wb/internal/infrastructure/repository/payment/model"
	"time"
)

type Order struct {
	OrderUid          string `json:"order_uid"`
	TrackNumber       string `json:"track_number"`
	Entry             string
	Delivery          repoDelivery.Delivery
	Payment           repoPayment.Payment
	Items             []repoItem.Item
	Locale            string
	InternalSignature string
	CustomerId        string
	DeliveryService   string
	Shardkey          string
	SmId              int
	DateCreated       time.Time
	OofShard          string
}
