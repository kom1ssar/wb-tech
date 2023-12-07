package order

import (
	"context"
)

func (s *service) GetByUUId(ctx context.Context, uuid string) {
	s.orderRepository.GetByUUId(ctx, uuid)
}
