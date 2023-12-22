package order

import (
	"sync"
	"tech-wb/internal/infrastructure/cache"
	"tech-wb/internal/model"
)

var _ cache.OrderCache = (*orderCache)(nil)

type orderCache struct {
	m    sync.RWMutex
	data map[string]*model.Order
}

func (c *orderCache) Set(key string, value *model.Order) error {
	c.m.Lock()
	defer c.m.Unlock()

	c.data[key] = value
	return nil
}

func (c *orderCache) Get(key string) *model.Order {
	c.m.RLock()
	defer c.m.RUnlock()

	order, ok := c.data[key]

	if !ok {
		return nil
	}

	return order
}

func NewOrderCache() cache.OrderCache {
	return &orderCache{
		m:    sync.RWMutex{},
		data: make(map[string]*model.Order),
	}
}
