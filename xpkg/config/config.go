package config

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	AppName string `env:"APP_NAME,default=football-manager-oss"`
	Env     string `env:"ENV,default=dev"`
	Host    string `env:"HOST,default=127.0.0.1"`
	Port    string `env:"PORT,default=8080"`
	Redis   *Redis
}

type Redis struct {
	Host     string `env:"REDIS_HOST,default=127.0.0.1"`
	Port     string `env:"REDIS_PORT,default=6379"`
	Password string `env:"REDIS_PASSWORD,default="`
}

func NewConfig() (*Config, error) {
	ctx := context.Background()
	config := new(Config)
	if err := envconfig.Process(ctx, config); err != nil {
		return nil, err
	}

	return config, nil
}
