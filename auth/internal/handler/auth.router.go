package handler

import "github.com/go-chi/chi"

func AuthRouter(r *chi.Mux, handler *AuthHandler) {

	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", handler.CreateUserHandler)
		r.Post("/sign-in", handler.LoginUserHandler)
		r.Get("/verify", VerifyTokenHandler)
		r.Get("/logout", LogoutSession)
	})
}
