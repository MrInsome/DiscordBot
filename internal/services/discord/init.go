package discord

import (
	"github.com/bwmarrin/discordgo"
)

func (cord *CordSession) InitCommands() error {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "hello",
			Description: "Say hello to the bot",
		},
		{
			Name:        "uvoice",
			Description: "Users",
		},
	}

	allcoms, _ := cord.ApplicationCommands(cord.State.User.ID, "")
	commap := make(map[string]string)
	for _, cmd := range allcoms {
		commap[cmd.Name] = cmd.ID
	}

	for _, command := range commands {
		if _, ok := commap[command.Name]; !ok {
			_, err := cord.ApplicationCommandCreate(cord.State.User.ID, "", command)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
