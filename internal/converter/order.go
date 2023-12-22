package converter

import (
	"tech-wb/internal/model"
	desc "tech-wb/pkg/gen/proto/order_v1"
)

func OrderToModelFromDesc(order *desc.Order) *model.Order {

	p := PaymentToModelFromDesc(order.Payment)
	d := DeliveryToModelFromDesc(order.Delivery)
	var items []model.Item

	for _, item := range order.Items {

		i := ItemToModelFromDesc(item)
		items = append(items, *i)
	}

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
		DateCreated:       order.DateCreated.AsTime(),
		OofShard:          order.OofShard,
	}
}
