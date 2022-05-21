package config

import (
	"os"
	"path/filepath"

	"github.com/spf13/viper"
)

const (
	PRODUCTION  = "production"
	STAGING     = "staging"
	DEVELOPMENT = "development"
)

func readConfig(conf *Schema, service string) error {
	viper.SetConfigFile(getConfigFile(service))
	return nil
}

func getConfigFile(service string) string {
	env := getEnv()
	basePath := "/config/environment/" + env
	return filepath.Join(basePath)
}

func getEnv() string {
	env := os.Getenv("ENV")

	switch env {
	case PRODUCTION:
		env = PRODUCTION
	case STAGING:
		env = STAGING
	default:
		env = DEVELOPMENT
	}

	return env
}
