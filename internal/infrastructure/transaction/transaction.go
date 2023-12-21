package transaction

import (
	"context"
	"tech-wb/internal/model"
)

type Transactions struct {
	OrderTransaction OrderTransaction
}

type OrderTransaction interface {
	Insert(ctx context.Context, order *model.Order) error
}
