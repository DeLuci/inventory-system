package main

import (
	"github.com/DeLuci/inventory-system/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func routes() http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Login)
	mux.Post("/", handlers.Repo.PostLogin)

	mux.Get("/signup", handlers.Repo.SignUp)
	mux.Post("/signup", handlers.Repo.PostSignUp)

	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer)) // helps use css/js

	mux.Route("/admin", func(mux chi.Router) {
		mux.Use(Auth)
		mux.Get("/dashboard", handlers.Repo.Dashboard)
		mux.Get("/search/{product}", handlers.Repo.SearchBarData)
		mux.Post("/scan", handlers.Repo.ScanProduct)
	})

	return mux
}
