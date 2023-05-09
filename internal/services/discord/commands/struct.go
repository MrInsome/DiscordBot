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
		"uvoice": c.UserVoice,
		"wakeup": c.WakeupCommand,

		"buttons": c.TestButtons, //todo
		"selects": c.Selects,     //todo
	}
	return commandsHandlers
}
