package services

import (
	"Discord_bot/internal/config"
	"Discord_bot/internal/contracts"
	"Discord_bot/internal/errors"
	"Discord_bot/internal/services/discord/commands"
	"Discord_bot/internal/services/discord/components"
	"github.com/bwmarrin/discordgo"
)

type Services struct {
	*discordgo.Session
	*config.Configs
	contracts.ErrorContract
	contracts.CommandsContract
	contracts.ComponentsContract
}

func NewServices(configs *config.Configs) (*Services, error) {
	dg, err := discordgo.New("Bot " + configs.Token)
	if err != nil {
		return nil, err
	}
	return &Services{
		Session:            dg,
		Configs:            configs,
		ErrorContract:      errors.NewDiscordErrors(dg),
		CommandsContract:   commands.NewCommands(dg),
		ComponentsContract: components.NewComponents(dg),
	}, nil
}
