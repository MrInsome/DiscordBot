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
}

func Init(cfg string) (*Configs, error) {
	err := godotenv.Load("./configs/" + cfg + ".env")
	if err != nil {
		return nil, err
	}

	token := os.Getenv("DISCORD_BOT_TOKEN")
	return &Configs{BotToken{Token: token}}, nil
}
