package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/prabowoteguh/belajar-vibe-code/config"
	"github.com/prabowoteguh/belajar-vibe-code/internal/handler"
	"github.com/prabowoteguh/belajar-vibe-code/internal/repository/sqlserver"
	"github.com/prabowoteguh/belajar-vibe-code/internal/service"
	"github.com/prabowoteguh/belajar-vibe-code/pkg/database"
	"github.com/prabowoteguh/belajar-vibe-code/pkg/logger"
	"github.com/prabowoteguh/belajar-vibe-code/pkg/redis"
	"github.com/prabowoteguh/belajar-vibe-code/routes"
	"go.uber.org/zap"
)

func main() {
	// Initialize Logger
	logger.InitLogger()
	defer func() {
		if err := logger.Log.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to sync logger: %v\n", err)
		}
	}()

	// Load Config
	cfg := config.LoadConfig()

	// Initialize Database
	db, err := database.InitSQLServer(cfg)
	if err != nil {
		logger.Fatal("failed to initialize sql server", zap.Error(err))
	}
	defer db.Close()

	// Initialize Redis
	_, err = redis.InitRedis(cfg)
	if err != nil {
		logger.Error("failed to initialize redis", zap.Error(err))
	} else {
		logger.Info("redis initialized successfully")
	}

	// Dependency Injection
	userRepo := sqlserver.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// Setup Router
	router := routes.SetupRoutes(userHandler)

	// Server Setup
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.AppPort),
		Handler:      router,
		ReadTimeout:  time.Duration(cfg.ReadTimeout) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeout) * time.Second,
		IdleTimeout:  time.Duration(cfg.IdleTimeout) * time.Second,
	}

	// Graceful Shutdown
	go func() {
		logger.Info("starting server", zap.String("port", cfg.AppPort))
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal("failed to start server", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logger.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("server forced to shutdown", zap.Error(err))
	}

	logger.Info("server exited properly")
}
