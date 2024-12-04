package config

import (
	"errors"
	"os"
)

type PgConnectionConfig struct {
	User          string
	Password      string
	Database      string
	Port          string
	ContainerName string
}

func LoadPgConfig() (PgConnectionConfig, error) {
	var config PgConnectionConfig

	config.User = os.Getenv("PG_USER")
	config.Password = os.Getenv("PG_PASSWORD")
	config.Database = os.Getenv("PG_DATABASE")
	config.Port = os.Getenv("PG_PORT")
	config.ContainerName = os.Getenv("PG_CONTAINER")

	if config.User == "" {
		return config, errors.New("no PG User")
	}
	if config.Password == "" {
		return config, errors.New("no PG Password")
	}
	if config.Database == "" {
		return config, errors.New("no PG Database")
	}
	if config.Port == "" {
		return config, errors.New("no PG Port")
	}
	if config.ContainerName == "" {
		return config, errors.New("no PG container name")
	}
	return config, nil
}
