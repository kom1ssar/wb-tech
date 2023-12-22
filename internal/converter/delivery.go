package converter

import (
	"tech-wb/internal/model"
	desc "tech-wb/pkg/gen/proto/order_v1"
)

func DeliveryToModelFromDesc(delivery *desc.Delivery) *model.Delivery {
	return &model.Delivery{
		Name:    delivery.Name,
		Phone:   delivery.Phone,
		Zip:     delivery.Zip,
		City:    delivery.City,
		Address: delivery.Address,
		Region:  delivery.Region,
		Email:   delivery.Email,
	}
}
