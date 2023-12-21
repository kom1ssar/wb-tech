package converter

import (
	"fmt"
	"tech-wb/internal/model"
	desc "tech-wb/pkg/gen/proto/order_v1"
	"time"
)

func OrderToModelFromDesc(order *desc.Order) *model.Order {

	p := PaymentToModelFromDesc(order.Payment)
	d := DeliveryToModelFromDesc(order.Delivery)
	var items []model.Item

	for _, item := range order.Items {

		i := ItemToModelFromDesc(item)
		items = append(items, *i)
	}

	fmt.Printf("%+v\n", items[0])
	fmt.Println("items")

	return &model.Order{
		OrderUid:          order.OrderUid,
		TrackNumber:       order.TrackNumber,
		Entry:             order.Entry,
		Delivery:          *d,
		Payment:           *p,
		Items:             items,
		Locale:            order.Locale,
		InternalSignature: order.InternalSignature,
		CustomerId:        order.CustomerId,
		DeliveryService:   order.DeliveryService,
		Shardkey:          order.ShardKey,
		SmId:              int(order.SmId),
		DateCreated:       time.Time{},
		OofShard:          order.OofShard,
	}
}
