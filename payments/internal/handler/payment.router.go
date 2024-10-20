package handler

import (
	"github.com/go-chi/chi"
)

func PaymentRouter(r *chi.Mux, handler *PaymentHandler) {

	r.Route("/payment", func(r chi.Router) {
		r.Post("/create", handler.CreateOrderHandler)
		r.Get("/success", handler.Success)
		r.Get("/failure", handler.Failure)
		r.Get("/pending", handler.Pending)
		r.Post("/webhook", handler.WebHook)
	})
}
