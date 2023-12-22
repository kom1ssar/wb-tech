package converter

import (
	"tech-wb/internal/model"
	desc "tech-wb/pkg/gen/proto/order_v1"
)

func PaymentToModelFromDesc(payment *desc.Payment) *model.Payment {

	return &model.Payment{
		RequestId:    payment.RequestId,
		Currency:     payment.Currency,
		Transaction:  payment.Transaction,
		Provider:     payment.Provider,
		Amount:       int(payment.Amount),
		PaymentDt:    int(payment.PaymentDt),
		Bank:         payment.Bank,
		DeliveryCost: int(payment.DeliveryCost),
		GoodsTotal:   int(payment.GoodsTotal),
		CustomFee:    int(payment.CustomFee),
	}
}
