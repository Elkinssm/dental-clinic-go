package config

import (
	"errors"
	"os"
)

type Config struct {
	PublicConfig  PublicConfig
	PrivateConfig PrivateConfig
}

type PublicConfig struct {
	PublicKey      string
	PostgresUser   string
	PostgresHost   string
	PostgresPort   string
	PostgresDBName string
}

type PrivateConfig struct {
	SecretKey        string
	PostgresPassword string
}

var (
	envs = map[string]PublicConfig{
		"local": {
			PublicKey:      "localAdmin",
			PostgresUser:   "elevenGroup",
			PostgresPort:   "5432",
			PostgresHost:   "localhost",
			PostgresDBName: "clinical_db",
		},
		"dev": {
			PublicKey: "dev_public_key",
		},
		"prod": {
			PublicKey: "prod_public_key",
		},
	}
)

func NewConfig(env string) (Config, error) {

	publicConfig, exists := envs[env]
	if !exists {
		return Config{}, errors.New("env not found")
	}

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		return Config{}, errors.New("secret key not found")
	}

	postgresPassword := os.Getenv("POSTGRES_PASSWORD")
	if postgresPassword == "" {
		return Config{}, errors.New("postgres password not found")
	}

	return Config{
		PublicConfig: publicConfig,
		PrivateConfig: PrivateConfig{
			SecretKey:        secretKey,
			PostgresPassword: postgresPassword,
		},
	}, nil
}
