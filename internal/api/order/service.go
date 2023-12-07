package order

import "tech-wb/internal/service"

type Implementation struct {
	orderService service.OrderService
}

func NewImplementation(orderService service.OrderService) *Implementation {
	return &Implementation{
		orderService: orderService,
	}
}
