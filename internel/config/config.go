package config

import (
	"os"
	"strings"

	"gorm.io/gorm"
)

type Config struct {
	Server ServerConfig
	App    AppConfig
}

type ServerConfig struct {
	Addr string
	DSN  string
	DB   *gorm.DB
}

type AppConfig struct {
	ENV     string
	Version string
}

func Load() Config {
	addr := os.Getenv("port")
	if addr != "" && !strings.Contains(addr, ":") {
		addr = ":" + addr
	}

	return Config{
		Server: ServerConfig{
			Addr: addr,
			DSN:  os.Getenv("dsn"),
		},
		App: AppConfig{
			ENV:     os.Getenv("env"),
			Version: os.Getenv("version"),
		},
	}
}
