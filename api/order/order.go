package order

import (
	"context"
	"github.com/go-chi/chi/v5"
	"tech-wb/internal/api/order"
)

func RegisterRoutes(ctx context.Context, router *chi.Mux, impl *order.Implementation) {

	router.Get("/order/{id}", impl.GetOrderById(ctx))

}
