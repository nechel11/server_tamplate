package apiserver

// Api server config
type Config struct{
	BindAddr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
}

//New config init
func NewConfig() *Config{
	return &Config{
		BindAddr: "8080",
		LogLevel: "debug",
	}
}