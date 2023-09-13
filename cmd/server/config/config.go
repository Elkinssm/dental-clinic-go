package config

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
			PublicKey:      "local_public_key",
			PostgresUser:   "eleven_group",
			PostgresPort:   "5432",
			PostgresHost:   "localhost",
			PostgresDBName: "eleven_group_clinic",
		},
		"dev": {
			PublicKey: "dev_public_key",
		},
		"prod": {
			PublicKey: "prod_public_key",
		},
	}
)
