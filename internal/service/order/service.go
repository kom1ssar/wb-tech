package order

import (
	"tech-wb/internal/infrastructure/repository"
	def "tech-wb/internal/service"
)

var _ def.OrderService = (*service)(nil)

type service struct {
	orderRepository repository.OrderRepository
}

func NewService(orderRepository repository.OrderRepository) *service {

	return &service{orderRepository: orderRepository}
}
