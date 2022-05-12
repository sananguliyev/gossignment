package config

import (
	"github.com/caarlos0/env/v6"
	"log"
)

type MongoDbConfig struct {
	Host     string `env:"MONGODB_HOST,notEmpty"`
	Username string `env:"MONGODB_USERNAME,notEmpty"`
	Password string `env:"MONGODB_PASSWORD,notEmpty"`
	Database string `env:"MONGODB_DATABASE,notEmpty"`
}

func NewMongoDbConfig() *MongoDbConfig {
	config := &MongoDbConfig{}
	if err := env.Parse(config); err != nil {
		log.Fatalf("Configuration error: %s", err)
	}

	return config
}
