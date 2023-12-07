package service

import "context"

type OrderService interface {
	GetByUUId(ctx context.Context, uuid string)
}
