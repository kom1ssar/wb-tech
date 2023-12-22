package order

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"net/http"
	"tech-wb/internal/api/response"
)

func (i *Implementation) GetOrderById(ctx context.Context) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		orderId := chi.URLParam(r, "id")

		o, err := i.orderService.GetByUUId(ctx, orderId)

		if err != nil {
			render.Status(r, 403)
			render.JSON(w, r, response.Error(err.Error(), 403))
			return
		}

		if o == nil {
			render.Status(r, 404)
			render.JSON(w, r, response.Error("order not found", 404))
			return
		}

		render.Status(r, 200)
		render.JSON(w, r, o)
		return

	}
}
