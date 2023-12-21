package app

import (
	"log"
	"tech-wb/internal/api/order"
	"tech-wb/internal/config"
	"tech-wb/internal/event"
	order_events "tech-wb/internal/event/order"
	"tech-wb/internal/infrastructure/repository"
	orderRepository "tech-wb/internal/infrastructure/repository/order"
	"tech-wb/internal/infrastructure/transaction"
	order_transaction "tech-wb/internal/infrastructure/transaction/order"
	"tech-wb/internal/service"
	orderService "tech-wb/internal/service/order"
	nats_streaming "tech-wb/pkg/client/nats-streaming"
	"tech-wb/pkg/client/postgresql"
)

type serviceProvider struct {
	httpConfig config.HTTPConfig

	cfgConfig config.DBConfig

	natsStreamingConfig config.NatsStreamConfig

	dbService postgresql.Client

	queueService nats_streaming.Client

	orderService service.OrderService

	orderRepository repository.OrderRepository

	orderImpl *order.Implementation

	orderSubscriptions event.OrderSubscriptions

	orderTransaction transaction.OrderTransaction
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

func (s *serviceProvider) NatsStreamingConfig() config.NatsStreamConfig {
	if s.natsStreamingConfig == nil {

		cfg, err := config.NewNatsStreamingConfig()
		if err != nil {
			log.Fatalf("failed to get Nats Streaming config %s", err.Error())
		}
		s.natsStreamingConfig = cfg
	}

	return s.natsStreamingConfig
}

func (s *serviceProvider) OrderRepository() repository.OrderRepository {
	if s.orderRepository == nil {
		s.orderRepository = orderRepository.NewRepository(s.dbService)
	}

	return s.orderRepository

}

func (s *serviceProvider) OrderService() service.OrderService {
	if s.orderService == nil {
		s.orderService = orderService.NewService(s.OrderRepository(), s.OrderTransaction())
	}
	return s.orderService
}

func (s *serviceProvider) OrderImpl() *order.Implementation {
	if s.orderImpl == nil {
		s.orderImpl = order.NewImplementation(s.OrderService())
	}

	return s.orderImpl

}

func (s serviceProvider) OrderSubscriptions() event.OrderSubscriptions {
	if s.orderSubscriptions == nil {
		s.orderSubscriptions = order_events.NewOrderSubscriptions(s.queueService, s.OrderService())
	}
	return s.orderSubscriptions
}

func (s *serviceProvider) OrderTransaction() transaction.OrderTransaction {
	if s.orderTransaction == nil {
		s.orderTransaction = order_transaction.NewOrderTransaction(s.dbService)
	}

	return s.orderTransaction
}
