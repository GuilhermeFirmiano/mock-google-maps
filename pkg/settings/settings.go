package settings

import (
	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

//Settings ...
type Settings struct {
	Host            string `env:"HOST"`
	ApplicationName string `env:"APPLICATION_NAME"`
}

// FromEnv ...
func FromEnv() *Settings {
	settings := new(Settings)

	err := env.Parse(settings)

	if err != nil {
		panic(err)
	}

	return settings
}

// FromDotEnv ...
func FromDotEnv(files ...string) *Settings {
	err := godotenv.Load(files...)

	if err != nil {
		panic(err)
	}

	return FromEnv()
}
