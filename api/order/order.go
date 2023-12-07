package order

import (
	"github.com/go-chi/chi/v5"
	"tech-wb/internal/api/order"
)

func RegisterUserRoutes(router *chi.Mux, impl *order.Implementation) {

	router.Get("/order/{id}", impl.GetOrderById())

}
