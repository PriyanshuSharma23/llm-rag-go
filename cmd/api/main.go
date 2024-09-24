package main

import "github.com/go-chi/chi/v5"

type application struct{}

func main() {
	app := &application{}
	router := chi.NewRouter()

	router.Post("/v1/upload", app.upload)
	router.Post("/v1/query", app.query)
}

func connectDB() {
}
