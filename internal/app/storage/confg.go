package storage

import (
	"fmt"
	"os"
)

// Config ...
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

// NewConfig ...
func NewConfig() *Config {
	return &Config{
		Host:     envLoader("host", "127.0.0.1"),
		Port:     envLoader("port", "5432"),
		User:     envLoader("user", "postgres"),
		Password: envLoader("password", "postgres"),
		DBName:   envLoader("dbname", "postgresdb"),
		SSLMode:  envLoader("sslmode", "disable"),
	}
}

// envLoader
func envLoader(name, defaultEnvString string) string {
	p := os.Getenv(name)
	if len(p) == 0 {
		return defaultEnvString
	}
	return p
}

// buildConnStr ...
func (c *Config) buildConnStr() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.Host,
		c.Port,
		c.User,
		c.Password,
		c.DBName,
		c.SSLMode,
	)
}
