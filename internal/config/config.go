package config

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/ilyakaznacheev/cleanenv"
)

// Config with tags from cleanenv library
type ServerConfig struct {
	Host                    string        `env:"API_HOST" env-required:"true"`
	Port                    string        `env:"API_PORT" env-required:"true"`
	MaxHeaderBytes          int           `yaml:"max_header_bytes" env-default:"1048576"`
	ReadTimeout             time.Duration `yaml:"read_timeout" env-default:"4s"`
	WriteTimeout            time.Duration `yaml:"write_timeout" env-default:"10s"`
	TimeForGracefulShutdown time.Duration `yaml:"time_for_graceful_shutdown" env-default:"10s"`
}

// Database config from env and config.yaml
type DatabaseConfig struct {
	Host     string `env:"DB_HOST" env-required:"true"`
	Port     string `env:"DB_PORT" env-required:"true"`
	Username string `env:"DB_USER" env-required:"true"`
	Password string `env:"DB_PASSWORD" env-required:"true"`
	DBName   string `env:"DB_NAME" env-required:"true"`
	SSLMode  string `env:"SSL_MODE" env-required:"true" validate:"oneof=disable require"`
}

// Logger config from config.yaml
type LoggerConfig struct {
	Level  string `yaml:"level" env-default:"info" validate:"oneof=debug info warn error"`
	Format string `yaml:"format" env-default:"json" validate:"oneof=text json"`
}

// Dataclass with all configs
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"-"`
	Logger   LoggerConfig   `yaml:"logger"`
}

// Load config from config/config.yaml
func LoadConfig(path string) (Config, error) {
	var config Config
	if err := cleanenv.ReadConfig(path, &config); err != nil {
		return Config{}, err
	}

	validate := validator.New()
	if err := validate.Struct(config); err != nil {
		return Config{}, err
	}

	return config, nil
}
