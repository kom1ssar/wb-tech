package converter

import (
	"tech-wb/internal/model"
	desc "tech-wb/pkg/gen/proto/order_v1"
	"time"
)

func PaymentToModelFromDesc(payment *desc.Payment) *model.Payment {

	return &model.Payment{
		RequestId:    payment.RequestId,
		Currency:     payment.Currency,
		Transaction:  payment.Transaction,
		Provider:     payment.Provider,
		Amount:       int(payment.Amount),
		PaymentDt:    time.Time{},
		Bank:         payment.Bank,
		DeliveryCost: int(payment.DeliveryCost),
		GoodsTotal:   int(payment.GoodsTotal),
		CustomFee:    int(payment.CustomFee),
	}
}
