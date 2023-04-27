package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func (cord *CordSession) InitCommands() {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "hello",
			Description: "Say hello to the bot",
		},
		{
			Name:        "voiceusers",
			Description: "Users",
		},
	}

	for _, command := range commands {
		_, err := cord.dg.ApplicationCommandCreate(cord.dg.State.User.ID, "", command)
		if err != nil {
			fmt.Println("Error registering command: ", err)
			return
		}
	}
}
