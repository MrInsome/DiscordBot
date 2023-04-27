package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

var VoiceUsers = make(map[string]bool)

func VoiceUsersFunc(dg *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "voiceusers" {
		data := discordgo.InteractionResponseData{}
		for userID := range VoiceUsers {
			user, err := dg.User(userID)
			if err != nil {
				log.Println(err)
				continue
			}
			avatarURL := user.AvatarURL("")

			data = discordgo.InteractionResponseData{
				Content: user.Username + " (" + user.ID + ")",
				Embeds: []*discordgo.MessageEmbed{
					{
						Image: &discordgo.MessageEmbedImage{
							URL: avatarURL,
						},
					},
				},
			}
		}
		dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &data,
		})
	}
}

func VoiceState(ds *discordgo.Session, vs *discordgo.VoiceStateUpdate) {
	if vs.GuildID != "1065951599891071046" {
		return
	}
	if vs.ChannelID == "1065951600541179988" {
		VoiceUsers[vs.UserID] = true
	} else {
		delete(VoiceUsers, vs.UserID)
	}
}
