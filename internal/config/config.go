package config

import (
	"github.com/caarlos0/env"
)

// nolint: lll
type Config struct {
	Port                 int    `env:"PORT" envDefault:"4000"`
	MySQLDatabase        string `env:"MYSQL_DATABASE" envDefault:"test_choice"`
	MySQLUser            string `env:"MYSQL_USER" envDefault:"test_user"`
	MySQLPassword        string `env:"MYSQL_PASSWORD" envDefault:"test_password"`
	MySQLRootPassword    string `env:"MYSQL_ROOT_PASSWORD" envDefault:"root_password"`
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
