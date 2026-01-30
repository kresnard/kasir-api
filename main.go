package main

import (
	"cashier_api/internal/category"
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	// middleware global
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		json.NewEncoder(w).Encode(map[string]string{
			"status":  "OK",
			"message": "service running",
		})
	})

	r.Route("/api", func(r chi.Router) {
		r.Route("/categories", func(r chi.Router) {
			r.Get("/", category.GetCategories)
			r.Get("/{id}", category.GetCategoryByID)
			r.Post("/", category.CreateCategory)
			r.Delete("/{id}", category.DeleteCategory)
			r.Put("/{id}", category.UpdateCategory)
		})
	})

	log.Println("server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
