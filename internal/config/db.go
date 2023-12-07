package config

import (
	"github.com/pkg/errors"
	"os"
	"strconv"
)

const (
	dbHost     = "DB_HOST"
	dbPort     = "DB_PORT"
	dbUser     = "DB_USER"
	dbPassword = "DB_PASSWORD"
	dbName     = "DB_NAME"
)

type DBConfig interface {
	GetHost() string
	GetPort() uint16
	GetUser() string
	GetPassword() string
	GetName() string
}

type dbConfig struct {
	host     string
	port     uint16
	user     string
	password string
	name     string
}

func NewDBConfig() (DBConfig, error) {
	host := os.Getenv(dbHost)
	if len(host) == 0 {
		return nil, errors.New("[.env] db host not found")
	}

	port := os.Getenv(dbPort)
	if len(port) == 0 {
		return nil, errors.New("[.env] db port not found")
	}

	user := os.Getenv(dbUser)
	if len(user) == 0 {
		return nil, errors.New("[.env] db user not found")
	}

	password := os.Getenv(dbPassword)
	if len(password) == 0 {
		return nil, errors.New("[.env] db password not found")
	}
	name := os.Getenv(dbName)
	if len(name) == 0 {
		return nil, errors.New("[.env] db name not found")
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return nil, errors.New("[.env] db port error parse to int")
	}

	return &dbConfig{
		host:     host,
		port:     uint16(portInt),
		user:     user,
		password: password,
		name:     name,
	}, nil
}

func (c *dbConfig) GetHost() string {
	return c.host

}

func (c *dbConfig) GetPort() uint16 {
	return c.port
}

func (c *dbConfig) GetUser() string {
	return c.user
}

func (c *dbConfig) GetPassword() string {
	return c.password
}

func (c *dbConfig) GetName() string {
	return c.name
}
