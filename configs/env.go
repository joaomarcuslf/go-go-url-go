package configs

import "os"

func FromEnv() (*Config, error) {
	cfg := Config{
		Server{
			Port: os.Getenv("PORT"),
		},
		Redis{
			Host:     os.Getenv("REDIS_HOST"),
			Port:     os.Getenv("REDIS_PORT"),
			Password: os.Getenv("REDIS_PASSWORD"),
		},
		Options{
			Schema: os.Getenv("SCHEMA"),
			Prefix: os.Getenv("PREFIX"),
			Mode:   os.Getenv("GIN_MODE"),
		},
	}

	return &cfg, nil
}
