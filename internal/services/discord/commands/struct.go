package commands

import (
	"Discord_bot/internal/contracts"
	"github.com/bwmarrin/discordgo"
)

type Commands struct {
	*discordgo.Session
	contracts.CommandsContract
}

func NewCommands(dg *discordgo.Session) *Commands {
	return &Commands{Session: dg}
}

func (c *Commands) RegisterCommands() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var commandsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"uvoice":  c.UserVoice,
		"buttons": c.TestButtons,
		"selects": c.Selects,
	}
	return commandsHandlers
}
