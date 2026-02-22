package config

import (
	"log"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	Env      Env    `env:"ENV" envDefault:"development"`
	Port     string `env:"PORT" envDefault:"2001"`
	SalonDsn string `env:"SALON_DSN,notEmpty"`
}

func New() Config {
	c, err := env.ParseAs[Config]()
	if err != nil {
		log.Fatal(err)
	}

	return c
}
