package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"time"
	"web_study/01/internal/store"
)

type application struct {
	config config
	store  store.Storage
}

type config struct {
	addr string
	db   dbConfig
}
type dbConfig struct {
	addr        string
	maxOpenCons int
	maxIdleCons int
	maxIdleTime time.Duration
}

func (app *application) mount() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	r.Route("/v1", func(r chi.Router) {
		r.Get("/health", app.healthHandler)

	})

	return r
}

func (app *application) run(mux *chi.Mux) error {
	srv := &http.Server{
		Addr:    app.config.addr,
		Handler: mux,
	}
	log.Println("Starting server on", app.config.addr)
	return srv.ListenAndServe()
}
