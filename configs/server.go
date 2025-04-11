package configs

import (
	"time"
)

type Configuration struct {
	HTTPServerConf
	LoggerConf
}

type HTTPServerConf struct {
	IdleTimeout  time.Duration `env:"HTTP_SERVER_IDLE_TIMEOUT" envDefault:"60s"`
	Port         int           `env:"PORT" envDefault:"8080"`
	ReadTimeout  time.Duration `env:"HTTP_SERVER_READ_TIMEOUT" envDefault:"1s"`
	WriteTimeout time.Duration `env:"HTTP_SERVER_WRITE_TIMEOUT" envDefault:"2s"`
}
