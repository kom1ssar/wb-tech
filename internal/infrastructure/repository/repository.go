package repository

import (
	"context"
	"tech-wb/internal/model"
)

type OrderRepository interface {
	GetByUUId(ctx context.Context, uuid string) (*model.Order, error)
}
