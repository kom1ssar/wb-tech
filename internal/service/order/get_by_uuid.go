package order

import (
	"context"
	"errors"
	"tech-wb/internal/model"
)

func (s *service) GetByUUId(ctx context.Context, uuid string) (*model.Order, error) {

	if len(uuid) <= 0 {
		return nil, errors.New("invalid order UUID")

	}

	o, err := s.orderRepository.GetByUUId(ctx, uuid)

	if err != nil {
		return nil, err
	}

	if o == nil || len(o.OrderUid) == 0 {
		return nil, nil
	}

	return o, nil
}
