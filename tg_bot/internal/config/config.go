package config

import (
	"fmt"
	"os"
	"sync"
)

type TConfig struct {
	TgBotToken string
	ServerPort string
}

var (
	once   sync.Once
	config *TConfig
)

func GetConfig() *TConfig {
	once.Do(func() {
		config = new(TConfig)
		config.TgBotToken = os.Getenv("TG_BOT_TOKEN")
		config.ServerPort = os.Getenv("TG_BOT_API_PORT")
		fmt.Println(config)
	})
	return config
}
