package config

import (
	"fmt"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Postgres
	Server

	ApiKey      string `env:"API_KEY" env-required:"true"`
	PathToPhoto string `env:"PATH_TO_PHOTO" env-required:"true"`
	PathToFont  string `env:"PATH_TO_FONT" env-required:"true"`
}

type Postgres struct {
	Uri string `env:"PG_URI" env-default:"postgresql://boombacks:blockme@localhost:5433/hackaton"`
}

type Server struct {
	Port string `env:"SRV_PORT" env-default:"3003"`
}

func NewConfig(configPath ...string) (*Config, error) {
	cfg := &Config{}
	var err error

	if len(configPath) == 0 {
		err = cleanenv.ReadEnv(cfg)
	} else {
		err = cleanenv.ReadConfig(filepath.Join("./", configPath[0]), cfg)
	}
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = cleanenv.UpdateEnv(cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %w", err)
	}

	return cfg, nil
}
