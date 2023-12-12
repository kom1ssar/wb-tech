package order

import (
	"context"
	"fmt"
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
			fmt.Println(err)
			render.JSON(w, r, response.Error(err.Error()))
			return
		}

		render.JSON(w, r, o)
		return

	}
}
