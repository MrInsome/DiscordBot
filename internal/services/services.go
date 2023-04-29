package services

import (
	"Discord_bot/internal/config"
	"Discord_bot/internal/contracts"
	"Discord_bot/internal/errors"
	"github.com/bwmarrin/discordgo"
)

type Services struct {
	*discordgo.Session
	*config.Configs
	contracts.ErrorContract
}

func NewServices(configs *config.Configs) (*Services, error) {
	dg, err := discordgo.New("Bot " + configs.Token)
	if err != nil {
		return nil, err
	}
	return &Services{
		Session:       dg,
		Configs:       configs,
		ErrorContract: errors.NewDiscordErrors(dg),
	}, nil
}
