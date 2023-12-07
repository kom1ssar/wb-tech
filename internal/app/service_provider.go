package app

import (
	"log"
	"tech-wb/internal/api/order"
	"tech-wb/internal/config"
	"tech-wb/internal/infrastructure/repository"
	orderRepository "tech-wb/internal/infrastructure/repository/order"
	"tech-wb/internal/service"
	orderService "tech-wb/internal/service/order"
	"tech-wb/pkg/client/postgresql"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig

	cfgConfig config.DBConfig

	dbService postgresql.Client

	orderService service.OrderService

	orderRepository repository.OrderRepository

	orderImpl *order.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPConfig() config.HTTPConfig {
	if s.httpConfig == nil {
		cfg, err := config.NewHTTPConfig()
		if err != nil {
			log.Fatalf("failed to get HTTP config: %s", err.Error())
		}

		s.httpConfig = cfg
	}

	return s.httpConfig
}

func (s *serviceProvider) DBConfig() config.DBConfig {
	if s.cfgConfig == nil {
		cfg, err := config.NewDBConfig()
		if err != nil {
			log.Fatalf("failed to get DB config: %s", err.Error())
		}
		s.cfgConfig = cfg
	}

	return s.cfgConfig

}

func (s *serviceProvider) OrderRepository() repository.OrderRepository {
	if s.orderRepository == nil {
		s.orderRepository = orderRepository.NewRepository(s.dbService)
	}

	return s.orderRepository

}

func (s *serviceProvider) OrderService() service.OrderService {
	if s.orderService == nil {
		s.orderService = orderService.NewService(s.OrderRepository())
	}
	return s.orderService
}

func (s *serviceProvider) OrderImpl() *order.Implementation {
	if s.orderImpl == nil {
		s.orderImpl = order.NewImplementation(s.OrderService())
	}

	return s.orderImpl

}
