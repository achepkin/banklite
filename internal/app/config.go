package app

import "os"

// Config is a config :).
type Config struct {
	HTTPAddr  string
	MaxAmount int
}

// Read reads config from environment.
func Read() Config {
	var config Config
	httpAddr, exists := os.LookupEnv("HTTP_ADDR")
	if exists {
		config.HTTPAddr = httpAddr
	} else {
		config.HTTPAddr = "127.0.0.1:8080"
	}

	return config
}
