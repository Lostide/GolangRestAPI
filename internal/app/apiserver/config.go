package apiserver

import "gitlab.alfa-bank.kz/alfamobile/credit/arrests-api/internal/app/store"

// Config describes server configuration
type Config struct {
	BinAddr  string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

func NewConfig() *Config {
	return &Config{
		BinAddr:  ":8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
