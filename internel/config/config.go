package config

import (
	"os"
	"strings"
)

type Config struct {
	Addr string
	DSN  string
}

func Load() Config {
	addr := os.Getenv("port")
	if addr != "" && !strings.Contains(addr, ":") {
		addr = ":" + addr
	}

	return Config{
		Addr: addr,
		DSN:  os.Getenv("dsn"),
	}
}
