package env

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

func Load(prefix string, out interface{}, filename ...string) error {
	err := godotenv.Load(filename...)
	if err != nil {
		return err
	}

	err = envconfig.Process(prefix, out)
	if err != nil {
		return err
	}
	return nil
}
