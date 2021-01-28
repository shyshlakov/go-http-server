package config

import (
	"fmt"
	"strings"
	"time"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	ServicePort             string        `env:"SERVICE_PORT,required"`
	PostgresHost            string        `env:"POSTGRES_HOST,required"`
	PostgresPort            string        `env:"POSTGRES_PORT,required" envDefault:"5432"`
	PostgresUsername        string        `env:"POSTGRES_USERNAME,required"`
	PostgresPassword        string        `env:"POSTGRES_PASSWORD,required"`
	PostgresDBName          string        `env:"POSTGRES_DBNAME,required"`
	PostgresSSLMode         string        `env:"POSTGRES_SSLMODE" envDefault:"disable"`
	PostgresConnTimeout     int           `env:"POSTGRES_CONN_TIMEOUT" envDefault:"5"`
	PostgresMaxOpenConns    int           `env:"POSTGRES_MAX_OPEN_CONNS" envDefault:"10"`
	PostgresMaxIdleConns    int           `env:"POSTGRES_MAX_IDLE_CONNS" envDefault:"3"`
	PostgresMaxConnLifetime time.Duration `env:"POSTGRES_MAX_CONN_LIFETIME" envDefault:"10s"`
}

func FromEnv() (*Config, error) {
	c := &Config{}
	if err := env.Parse(c); err != nil {
		return nil, err
	}
	return c, nil
}
func (c *Config) GetDSN(driver string) (string, error) {
	switch strings.ToLower(driver) {
	case "postgres":
		return c.getPostgresDSN(), nil
	default:
		return "", fmt.Errorf("unknown driver %q", driver)
	}
}

// GetPostgresDSN generates postgres dsn string
func (c *Config) getPostgresDSN() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.PostgresUsername,
		c.PostgresPassword,
		c.PostgresHost,
		c.PostgresPort,
		c.PostgresDBName,
		c.PostgresSSLMode)
}
