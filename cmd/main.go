package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"go-crud-api/internal/handler"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/todos", handler.RegisterTodoRoutes)

	http.ListenAndServe(":8080", r)
}
