package order

import (
	"tech-wb/internal/infrastructure/repository"
	"tech-wb/internal/infrastructure/transaction"
	def "tech-wb/internal/service"
)

var _ def.OrderService = (*service)(nil)

type service struct {
	orderRepository  repository.OrderRepository
	orderTransaction transaction.OrderTransaction
}

func NewService(orderRepository repository.OrderRepository, orderTransaction transaction.OrderTransaction) def.OrderService {
	return &service{
		orderRepository:  orderRepository,
		orderTransaction: orderTransaction,
	}
}
