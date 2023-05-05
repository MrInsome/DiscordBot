package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Configs struct {
	BotToken
}

type BotToken struct {
	Token string
	AppID string
}

func Init(cfg string) (*Configs, error) {
	err := godotenv.Load("./configs/" + cfg + ".env")
	if err != nil {
		return nil, err
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	appID := os.Getenv("APP_ID")
	return &Configs{BotToken{Token: token, AppID: appID}}, nil
}
