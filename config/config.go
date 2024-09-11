package config

import (
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	App
	Http
	Database
}

type Http struct {
	Port string
}

type App struct {
	Name    string
	Version string
}

type Database struct {
	PoolSize     int
	DATABASE_URL string
}

func DefineConfig() (*Config, error) {
	env, err := godotenv.Read()
	if err != nil {
		return nil, err
	}

	config := &Config{
		App: App{
			Name:    env["APP_NAME"],
			Version: env["APP_VERSION"],
		},
		Http: Http{
			Port: env["HTTP_PORT"],
		},
		Database: Database{
			DATABASE_URL: env["DATABASE_URL"],
		},
	}
	pz := env["DATABASE_POOL_SIZE"]
	config.Database.PoolSize, _ = strconv.Atoi(pz)
	return config, nil
}
