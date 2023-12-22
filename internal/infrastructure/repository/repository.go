package repository

import (
	"context"
	"tech-wb/internal/model"
)

type OrderRepository interface {
	GetByUUId(ctx context.Context, uuid string) (*model.Order, error)
	GetListDESCCreated(ctx context.Context, limit int) ([]*model.Order, error)
}
