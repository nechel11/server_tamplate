package apiserver

import "github.com/nechel11/Ozon_test/internal/app/store"

// Api server config
type Config struct{
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store *store.Config
}

//New config init
func NewConfig() *Config{
	return &Config{
		BindAddr: ":8080",
		LogLevel: "debug",
		Store: store.NewConfig(),
	}
}