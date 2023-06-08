package envx

import (
	"github.com/caarlos0/env/v6"
	"github.com/labstack/gommon/log"
)

func LoadEnv(v interface{}, opts ...env.Options) error {
	return env.Parse(v, opts...)
}

func MustLoadEnv(v interface{}, opts ...env.Options) {
	if err := LoadEnv(v, opts...); err != nil {
		log.Fatal(err)
	}
}
