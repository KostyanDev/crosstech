package config

import (
	"time"
)

type Config struct {
	App        App
	HTTPServer HTTPServer `envPrefix:"HTTP_"`
	Storage    Storage    `envPrefix:"STORAGE_"`
}

type App struct {
	Name string `env:"APP_NAME" envDefault:"app"`
}

type HTTPServer struct {
	Host         string        `env:"SERVER_HOST" envDefault:"localhost"`
	Port         int           `env:"SERVER_PORT" envDefault:"8080"`
	WriteTimeout time.Duration `env:"WRITE_TIMEOUT" envDefault:"15s"`
	ReadTimeout  time.Duration `env:"READ_TIMEOUT" envDefault:"15s"`
}

type Storage struct {
	DSN     string        `env:"DSN" valid:"required" envDefault:"postgres://myuser:mypassword@postgres:5432/mydatabase?sslmode=disable"`
	Timeout time.Duration `env:"QUERY_TIMEOUT" envDefault:"5s"`
}
