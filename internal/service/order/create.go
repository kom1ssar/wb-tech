package order

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"tech-wb/internal/model"
)

func (s *service) Create(ctx context.Context, order *model.Order) error {

	already, err := s.orderRepository.GetByUUId(ctx, order.OrderUid)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(already.OrderUid) != 0 {
		return errors.New("orderId already exists")
	}

	err = s.orderTransaction.Insert(ctx, order)
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil

}
