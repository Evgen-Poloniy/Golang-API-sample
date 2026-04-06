package app

import (
	"api_sample/internal/config"
	"api_sample/internal/repository"
	httpserver "api_sample/internal/server/http"
	"api_sample/internal/service"
	"api_sample/internal/transport/http/base_url"
	v1 "api_sample/internal/transport/http/v1"
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"

	// "api_sample/pkg/database"
	"api_sample/pkg/logger"
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
func Run() {
	apiKey := os.Getenv("API_KEY")
	if apiKey == "" {
		logrus.Fatal("required API_KEY")
	}

	apiKeyHash, err := bcrypt.GenerateFromPassword([]byte(apiKey), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatalf("error when generation API_KEY hash: %v", err)
	}

	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		logrus.Fatalf("error when loading env CONFIG_PATH")
	}

	config, err := config.LoadConfig(configPath)
	if err != nil {
		logrus.Fatalf("error when loading config: %v", err)
	}

	loggerConfig := logger.Config{
		Level:  config.Logger.Level,
		Format: config.Logger.Format,
		Output: config.Logger.Output,
		Files:  config.Logger.Files,
	}

	logger, err := logger.NewLogrusLogger(&loggerConfig)
	if err != nil {
		logrus.Fatalf("error when loading the logger: %v", err)
	}
	defer func() {
		if err := logger.Close(); err != nil {
			logger.Errorf("error when closing the logger: %v", err)
		}
	}()

	// db, err := database.NewPostgreSQL(&config.Database)
	// if err != nil {
	// 	logger.Fatal("the database initialize error: %v", err)
	// }
	// defer func() {
	// 	if err := db.Close(); err != nil {
	// 		logger.Errorf("error when closing connection with the database: %v", err)
	// 		return
	// 	}
	// 	logger.Info("database connection has been closed")
	// }()
	// logger.Info("database connection established successful")

	// gatewat := gateway.NewHTTPGateway(&http.Client{})
	repository := repository.NewRepository(nil /*db*/, logger)
	service := service.NewService(repository, logger)
	baseHandler := base_url.NewHandler(service, &config.Info)
	v1Handler := v1.NewHandler(service)

	router := base_url.NewRouter(baseHandler, logger)
	v1.NewRouter(router, v1Handler, apiKeyHash)

	var wg sync.WaitGroup

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(1)
	server := httpserver.NewServer(&config.Server, router)
	go func() {
		defer wg.Done()

		logger.Infof("server is running on %s:%s", config.Server.Host, config.Server.Port)
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

	logger.Info("shutting down server")
	if err := server.Shutdown(shutdownCtx); err != nil {
		logger.Errorf("server forced to shutdown: %v", err)
	}

	wg.Wait()

	logger.Info("server stop")
}
