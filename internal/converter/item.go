package converter

import (
	"tech-wb/internal/model"
	desc "tech-wb/pkg/gen/proto/order_v1"
)

func ItemToModelFromDesc(item *desc.Item) *model.Item {

	return &model.Item{
		ChrtId:      int(item.ChrtId),
		TrackNumber: item.TrackNumber,
		Price:       int(item.Price),
		Rid:         item.Rid,
		Name:        item.Name,
		Sale:        int(item.Sale),
		Size:        item.Size,
		TotalPrice:  int(item.TotalPrice),
		NmId:        int(item.NmId),
		Brand:       item.Brand,
		Status:      int(item.Status),
	}
}
