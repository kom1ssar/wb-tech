package nats_streaming

import (
	"context"
	"fmt"
	"github.com/nats-io/stan.go"
	"log"
	"tech-wb/internal/config"
)

type Client interface {
	Publish(subject string, data []byte) error
	Subscribe(subject string, cb stan.MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error)
	QueueSubscribe(subject, qgroup string, cb stan.MsgHandler, opts ...stan.SubscriptionOption) (stan.Subscription, error)
}

func NewClient(_ context.Context, config config.NatsStreamConfig) (*stan.Conn, error) {

	sc, err := stan.Connect(
		config.GetClusterId(), config.GetClientId(),
		stan.Pings(1, 3), stan.NatsURL(config.GetURL()),
	)

	if err != nil {
		fmt.Println("Nats Streaming connection err")
		log.Fatalln(err)
		return nil, err
	}

	return &sc, nil
}
