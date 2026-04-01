package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"project/internal/config"
	"project/internal/database"
	"project/internal/handler"
	"project/internal/logutil"
	"project/internal/repository"
	"project/internal/server"
	"project/internal/service"
	"sync"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	// _ "project/docs"
)

// @title File Parser API
// @version 1.0
// @description API for parsing files and retrieving parsed data and errors.
// @host localhost:3505
// @BasePath /api/v1
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		log.Fatalln("error when loading env API_KEY")
	}

	apiKeyHash, err := bcrypt.GenerateFromPassword([]byte(apiKey), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("error when generation API_KEY hash: %v", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		log.Fatalln("error when loading env CONFIG_PATH")
	}

	config, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("error when loading config: %v", err)
	}

	logger := logrus.New()
	logger.SetOutput(os.Stdout)
	logutil.SetLevel(logger, &config.Logger)
	logutil.SetFormatter(logger, &config.Logger)

	db, err := database.NewPostgreSQLDatabase(&config.Database)
	if err != nil {
		logger.Fatalf("database initialize error: %v", err)
	}
	defer func() {
		db.Close()
		logger.Println("database connection has been closed")
	}()
	logger.Println("database connection established successful")

	repository := repository.NewRepository(db)
	service := service.NewService(repository)
	handler := handler.NewHandler(service)

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	server := server.NewServer(&config.Server, handler.InitHandlers(&config.Logger, apiKeyHash))
	go func() {
		defer wg.Done()

		logger.Printf("server is running on %s:%s", config.Server.Host, config.Server.Port)
		if err := server.Start(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Errorf("server error: %v", err)
			cancel()
		}
	}()

	// jobs := make(chan string, 100)
	// defer close(jobs)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	select {
	case <-quit:
		cancel()
	case <-ctx.Done():
	}

	shutdownCtx, shutdownCancel := context.WithTimeout(
		context.Background(),
		config.Server.TimeForGracefulShutdown*time.Second,
	)
	defer shutdownCancel()

	logger.Println("shutting down server")
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Fatalf("server forced to shutdown: %v", err)
	}

	wg.Wait()

	logger.Println("server exited properly")
}
