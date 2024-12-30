package handlers

import (
	"github.com/gabriel-q7/go-rest-api/internal/middleware"
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
)

func Handler(r *chi.Mux) {
	// GLobal middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {

		router.Use(middleware.Authorization)

		router.Get("/coins", GetCoinBalance)
	})
}
