package order

import (
	"context"
	"github.com/pkg/errors"
	"tech-wb/internal/model"
	"tech-wb/pkg/utils/validation"
)

func (s *service) GetByUUId(ctx context.Context, uuid string) (*model.Order, error) {
	if !validation.UUID(uuid) {

		return nil, errors.New("Invalid order UUID")
	}

	order, err := s.orderRepository.GetByUUId(ctx, uuid)

	if err != nil {
		return nil, err
	}

	return order, nil
}
