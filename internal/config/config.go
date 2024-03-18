package config

import (
	"github.com/caarlos0/env"
)

// nolint: lll
type Config struct {
	Port                 int    `env:"PORT" envDefault:"4000"`
	MySQLConnect         string `env:"MYSQL_CONNECTION" envDefault:"root:district13@tcp(10.52.1.60:3306)/best_option?parseTime=true"`
	MySQLMaxConnections  int    `env:"MYSQL_MAX_CONNECTIONS" envDefault:"10"`
	MySQLIdleConnections int    `env:"MYSQL_IDLE_CONNECTIONS" envDefault:"2"`
	Level                string `env:"LOG_LEVEL" envDefault:"info"`
}

func New() (*Config, error) {
	var cfg Config

	if err := env.Parse(&cfg); err != nil {
		return &cfg, err
	}

	return &cfg, nil
}
