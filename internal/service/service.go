package service

import (
	"context"
	"tech-wb/internal/model"
)

type OrderService interface {
	GetByUUId(ctx context.Context, uuid string) (*model.Order, error)
}
