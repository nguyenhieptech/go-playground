package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/nguyenhieptech/go-playground/chi-apis/internal/handlers"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRoutes(collection *mongo.Collection) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/api/todos", handlers.GetTodos(collection))
	r.Post("/api/todos", handlers.CreateTodo(collection))
	r.Patch("/api/todos/{id}", handlers.UpdateTodo(collection))
	r.Delete("/api/todos/{id}", handlers.DeleteTodo(collection))

	return r
}
