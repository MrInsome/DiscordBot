package discord

import (
	"github.com/bwmarrin/discordgo"
)

func (cord *CordSession) InitCom(name string, description string, guildId string) error {
	var command discordgo.ApplicationCommand
	command.Name = name
	command.Description = description
	_, err := cord.ApplicationCommandCreate(cord.State.User.ID, guildId, &command)
	if err != nil {
		return err
	}
	return nil
}

func (cord *CordSession) DeleteCom(name string, guildId string) error {
	allcoms, _ := cord.ApplicationCommands(cord.State.User.ID, guildId)
	for _, cmd := range allcoms {
		if cmd.Name == name {
			err := cord.ApplicationCommandDelete(cord.State.User.ID, guildId, cmd.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//func (cord *CordSession) InitCommands() error {
//	commands := []*discordgo.ApplicationCommand{
//		{
//			Name:        "uvoice",
//			Description: "Users",
//		},
//	}
//
//	allcoms, _ := cord.ApplicationCommands(cord.State.User.ID, "")
//	commap := make(map[string]string)
//	for _, cmd := range allcoms {
//		commap[cmd.Name] = cmd.ID
//	}
//
//	for _, command := range commands {
//		if _, ok := commap[command.Name]; !ok {
//			_, err := cord.ApplicationCommandCreate(cord.State.User.ID, "", command)
//			if err != nil {
//				return err
//			}
//		}
//	}
//
//	return nil
//}
