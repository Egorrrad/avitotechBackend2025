package configs

import (
	"github.com/caarlos0/env/v9"
)

func Load() (*Configuration, error) {
	cfg := Configuration{}

	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
