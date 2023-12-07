package order

import (
	"context"
	def "tech-wb/internal/infrastructure/repository"
	"tech-wb/internal/model"
	"tech-wb/pkg/client/postgresql"
)

var _ def.OrderRepository = (*repository)(nil)

type repository struct {
	db postgresql.Client
}

func (r *repository) GetByUUId(ctx context.Context, uuid string) (*model.Order, error) {
	panic("implement me")
}

func NewRepository(db postgresql.Client) *repository {

	return &repository{db: db}

}
