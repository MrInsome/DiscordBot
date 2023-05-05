package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var VoiceUsers = make(map[string]bool)

func (cord *CordSession) Ready(dg *discordgo.Session, event *discordgo.Ready) {
	_, err := cord.ChannelMessageSend("1065951600541179987", "Бот запущен")
	if err != nil {
		log.Printf("Error from Ready: %s", err)
	}
}

func (cord *CordSession) VoiceState(ds *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	if vs.GuildID != "1065951599891071046" {
		return
	}
	if vs.ChannelID == "1065951600541179988" {
		VoiceUsers[vs.UserID] = true
	} else {
		delete(VoiceUsers, vs.UserID)
	}
}

func (cord *CordSession) ApplicationOrMessageComponent(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		commandsHandlers := cord.RegisterCommands()
		if h, ok := commandsHandlers[i.ApplicationCommandData().Name]; ok {
			h(s, i)
		}
	case discordgo.InteractionMessageComponent:
		componentsHandlers := cord.RegisterComponents()
		if h, ok := componentsHandlers[i.MessageComponentData().CustomID]; ok {
			h(s, i)
		}
	}
}
