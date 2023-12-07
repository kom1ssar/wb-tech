package order

import (
	"context"
	def "tech-wb/internal/infrastructure/repository"
	"tech-wb/internal/infrastructure/storage"
	"tech-wb/internal/model"
)

var _ def.OrderRepository = (*repository)(nil)

type repository struct {
	db *storage.Postgres
}

func (r *repository) GetByUUId(ctx context.Context, uuid string) (*model.Order, error) {
	panic("implement me")
}

func NewRepository(db *storage.Postgres) *repository {

	return &repository{db: db}

}
