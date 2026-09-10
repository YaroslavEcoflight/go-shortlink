package config

import (
	"fmt"

	"github.com/caarlos0/env/v11"
)

type (
	Config struct {
		App       app       `envPrefix:"APP_"`
		HTTP      http
		Secret    secret
		Pg        pg        `envPrefix:"POSTGRES_"`
		Redis     redis     `envPrefix:"REDIS_"`
		Analytics analytics `envPrefix:"ANALYTICS_"`
	}

	app struct {
		Name     string `env:"NAME,required"`
		Version  string `env:"VERSION,required"`
		LogLevel string `env:"LOG_LEVEL" envDefault:"info"`
	}

	secret struct {
		JWTSecret string `env:"JWT_SECRET,required"`
	}

	http struct {
		Port string `env:"HTTP_PORT,required"`
	}

	pg struct {
		Host     string `env:"HOST,required"`
		Port     int    `env:"PORT,required"`
		User     string `env:"USER,required"`
		Password string `env:"PASSWORD,required"`
		DBName   string `env:"DBNAME,required"`
		SSLMode  string `env:"SSLMODE,required"`
	}

	redis struct {
		Addr     string `env:"ADDR,required"`
		User     string `env:"HOST,required"`
		Password string `env:"PORT,required"`
	}

	analytics struct {
		Addr string `env:"ADDR,required"`
	}
)

func (p pg) DSN() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		p.Host, p.Port, p.User, p.Password, p.DBName, p.SSLMode)
}

func (r redis) URL() string {
	return r.Addr
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, fmt.Errorf("config error: %w", err)
	}
	return cfg, nil
}
