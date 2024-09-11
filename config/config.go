package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Http `yaml:"http"`
}

type Http struct {
	Port string `env-required:"true" yaml:"port" env:"HTTP_PORT"`
}

type App struct {
	Name    string `env-required:"true" yaml:"name" env:"APP_NAME"`
	Version string `env-required:"true" yaml:"version" env:"APP_VERSION"`
}

type Database struct{}

func DefineConfig() (*Config, error) {
	config := &Config{}

	err := cleanenv.ReadConfig("./config/config.yml", config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
