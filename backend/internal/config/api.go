package config

import (
	"errors"
	"log"
	"os"
	"strconv"
)

type ApiConfig struct {
	Port   int
	Secret string
}

var conf *ApiConfig

func LoadApiConfig() (ApiConfig, error) {
	conf = new(ApiConfig)
	var err error
	portString := os.Getenv("BACKEND_PORT")
	if portString == "" {
		return *conf, errors.New("BACKEND_PORT variable not set")
	}
	conf.Secret = os.Getenv("JWT_SECRET")
	log.Println(conf.Secret)
	conf.Port, err = strconv.Atoi(portString)
	return *conf, err
}

func GetJwtSecret() string {
	return conf.Secret
}
