package config

import (
	"net"
	"os"
	"time"

	"github.com/pkg/errors"
)

const (
	httpHostEnvName        = "HTTP_HOST"
	httpPortEnvName        = "HTTP_PORT"
	httpTimeoutEnvName     = "HTTP_TIMEOUT"
	httpIdleTimeoutEnvName = "HTTP_IDLE_TIMEOUT"
)

type HTTPConfig interface {
	Address() string
	GetTimeout() time.Duration
	GetIdleTimeout() time.Duration
}

type httpConfig struct {
	host        string
	port        string
	timeout     time.Duration
	idleTimeout time.Duration
}

func NewHTTPConfig() (HTTPConfig, error) {
	host := os.Getenv(httpHostEnvName)
	if len(host) == 0 {
		return nil, errors.New("[.env] http host not found")
	}

	port := os.Getenv(httpPortEnvName)
	if len(port) == 0 {
		return nil, errors.New("[.env] http port not found")
	}

	return &httpConfig{
		host:        host,
		port:        port,
		timeout:     4 * time.Second,  //todo
		idleTimeout: 30 * time.Second, //todo
	}, nil
}

func (c *httpConfig) Address() string {
	return net.JoinHostPort(c.host, c.port)
}

func (c *httpConfig) GetTimeout() time.Duration {
	return c.timeout
}

func (c *httpConfig) GetIdleTimeout() time.Duration {
	return c.idleTimeout
}
