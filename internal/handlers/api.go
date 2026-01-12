package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"github.com/dexterhere04/Learn_go/internal/middleware"
)

func Handler(r *chi.Mux){
	r.use(chimiddle.StripSlashes)
	r.Route("/account",func(router chi.Router)){
		router.Use(middleware.Authorization)
	}
}