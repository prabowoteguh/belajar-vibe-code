package routes

import (
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/prabowoteguh/belajar-vibe-code/internal/handler"
	"github.com/prabowoteguh/belajar-vibe-code/internal/middleware"
)

func SetupRoutes(userHandler *handler.UserHandler) *chi.Mux {
	r := chi.NewRouter()

	// Standard middlewares
	r.Use(chiMiddleware.RequestID)
	r.Use(chiMiddleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recovery)
	r.Use(middleware.Timeout(15 * time.Second))

	// Health check
	r.Get("/health", userHandler.HealthCheck)

	// User routes
	r.Route("/users", func(r chi.Router) {
		r.Get("/", userHandler.GetUsers)
		r.Post("/", userHandler.CreateUser)
	})

	return r
}
