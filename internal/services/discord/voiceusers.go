package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"time"
)

var VoiceUsers = make(map[string]bool)

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

func VoiceUsersFunc(dg *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "uvoice" {
		channel, err := dg.State.Channel(i.ChannelID)
		if err != nil {
			log.Printf("Ошибка получения канала от DG : %s", err)
			return
		}

		guild, err := dg.State.Guild(channel.GuildID)
		if err != nil {
			log.Printf("Ошибка получения сервера от DG : %s", err)
			return
		}

		usersInVoice := make(map[string][]string)

		for _, vs := range guild.VoiceStates {
			if vs.ChannelID != "" {
				if _, ok := usersInVoice[vs.ChannelID]; !ok {
					usersInVoice[vs.ChannelID] = []string{}
				}
				usersInVoice[vs.ChannelID] = append(usersInVoice[vs.ChannelID], vs.UserID)
			}
		}
		data := discordgo.InteractionResponseData{}
		for channelID, users := range usersInVoice {
			channel, err = dg.Channel(channelID)
			if err != nil {
				log.Printf("Ошибка получения войс канала от DG : %s", err)
				continue
			}

			for _, userID := range users {
				user, err := dg.User(userID)
				if err != nil {
					log.Printf("Ошибка получения ID пользователя от DG : %s", err)
					continue
				}

				rand.NewSource(time.Now().UnixNano())
				color := rand.Intn(16777215)

				avatarURL := user.AvatarURL("")
				embed := &discordgo.MessageEmbed{Color: color, Title: user.Username + "\n" + user.ID, Image: &discordgo.MessageEmbedImage{
					URL: avatarURL,
				}}

				dm, err := dg.UserChannelCreate(i.Interaction.Member.User.ID)
				if err != nil {
					log.Printf("Ошибка создания приватного сообщения : %s", err)
					continue
				}

				go dg.ChannelMessageSendEmbed(dm.ID, embed)
			}

			data = discordgo.InteractionResponseData{
				Content: "Команда выполнена",
				Flags:   discordgo.MessageFlagsEphemeral,
			}
			dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
				Type: discordgo.InteractionResponseChannelMessageWithSource,
				Data: &data,
			})
		}
	}
}
