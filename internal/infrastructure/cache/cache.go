package cache

import "tech-wb/internal/model"

type OrderCache interface {
	Set(key string, value *model.Order) error
	Get(key string) (*model.Order, bool)
}
