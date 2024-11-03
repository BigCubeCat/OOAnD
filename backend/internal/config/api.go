package config

import (
	"errors"
	"os"
	"strconv"
)

type ApiConfig struct {
	Port   int
	Secret string
}

var conf ApiConfig

func LoadApiConfig() (ApiConfig, error) {
	var cfg ApiConfig
	var err error
	portString := os.Getenv("BACKEND_PORT")
	if portString == "" {
		return cfg, errors.New("BACKEND_PORT variable not set")
	}
	cfg.Secret = os.Getenv("JWT_SECRET")
	cfg.Port, err = strconv.Atoi(portString)
	if err != nil {
		conf = cfg
	}
	return cfg, err
}

func GetJwtSecret() string {
	return conf.Secret
}
