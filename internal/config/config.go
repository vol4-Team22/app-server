package config

import "github.com/caarlos0/env/v6"

type Config struct {
	Env        string `env:"MIKKE_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"80"`
	DBHost     string `env:"MIKKE_DB_HOST" envDefault:"127.0.0.1"`
	DBPort     int    `env:"MIKKE_DB_PORT" envDefault:"33306"`
	DBUser     string `env:"MIKKE_DB_USER" envDefault:"user"`
	DBPassword string `env:"TMIKKE_DB_PASSWORD" envDefault:"password"`
	DBName     string `env:"MIKKE_DB_NAME" envDefault:"mikke"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
