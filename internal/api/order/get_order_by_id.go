package order

import (
	"net/http"
)

func (i *Implementation) GetOrderById() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
	}
}
