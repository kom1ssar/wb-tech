package config

import (
	"github.com/pkg/errors"
	"os"
)

const (
	nSClusterId = "NATS_CLUSTER_ID"
	nSClientId  = "NATS_CLIENT_ID"
	nSURL       = "NATS_URL"
)

type NatsStreamConfig interface {
	GetClusterId() string
	GetClientId() string
	GetURL() string
}

type natsStreamingConfig struct {
	clusterId string
	clientId  string
	url       string
}

func NewNatsStreamingConfig() (NatsStreamConfig, error) {
	clusterId := os.Getenv(nSClusterId)
	if len(clusterId) == 0 {
		return nil, errors.New("[.env] nats streaming clusterId not found")
	}

	clientId := os.Getenv(nSClientId)
	if len(clientId) == 0 {
		return nil, errors.New("[.env] nats streaming clusterId not found")
	}

	url := os.Getenv(nSURL)
	if len(url) == 0 {
		return nil, errors.New("[.env] nats streaming clusterId not found")
	}

	return &natsStreamingConfig{
		clusterId: clusterId,
		clientId:  clientId,
		url:       url,
	}, nil

}

func (c *natsStreamingConfig) GetClusterId() string {
	return c.clusterId
}

func (c *natsStreamingConfig) GetClientId() string {
	return c.clientId
}

func (c *natsStreamingConfig) GetURL() string {
	return c.url
}
