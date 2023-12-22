package order

import (
	"context"
	"fmt"
	"github.com/nats-io/stan.go"
	"google.golang.org/protobuf/proto"
	"log"
	"tech-wb/internal/converter"
	def "tech-wb/internal/event"
	"tech-wb/internal/service"
	"tech-wb/pkg/client/nats-streaming"
	desc "tech-wb/pkg/gen/proto/order_v1"
)

var _ def.OrderSubscriptions = (*subscriptions)(nil)

type subscriptions struct {
	queueService nats_streaming.Client
	orderService service.OrderService
}

func (s *subscriptions) Subscribe(ctx context.Context) {
	s.orderNew(ctx)
}

func (s *subscriptions) orderNew(ctx context.Context) {
	subject := "order.new"
	orderProto := desc.Order{}

	s.queueService.Subscribe(subject, func(msg *stan.Msg) {

		if err := proto.Unmarshal(msg.Data, &orderProto); err != nil {
			log.Printf("Failed to unmarshal message: %v", err)
			fmt.Printf(err.Error())
			return
		}

		order := converter.OrderToModelFromDesc(&orderProto)

		err := s.orderService.Create(ctx, order)
		if err != nil {
			fmt.Println(err.Error())
		}

	})

}

func NewOrderSubscriptions(queueService nats_streaming.Client, orderService service.OrderService) *subscriptions {
	return &subscriptions{
		queueService: queueService,
		orderService: orderService,
	}

}
