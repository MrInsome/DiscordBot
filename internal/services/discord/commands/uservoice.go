package commands

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"math/rand"
	"time"
)

func (c *Commands) UserVoice(s *discordgo.Session, i *discordgo.InteractionCreate) {
	channel, err := s.State.Channel(i.ChannelID)
	if err != nil {
		//cord.ErrorContract.ErrorMessage(err, i)
		log.Printf("Error from GetGuild : %s", err)
		return
	}

	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		//cord.ErrorContract.ErrorMessage(err, i)
		log.Printf("Error from cord GetGuild : %s", err)
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
		channel, err = s.Channel(channelID)
		if err != nil {
			log.Printf("Error from GetChanel ID : %s", err)
			continue
		}

		for _, userID := range users {
			user, err := s.User(userID)
			if err != nil {
				log.Printf("Error from GetUserID : %s", err)
				continue
			}

			rand.NewSource(time.Now().UnixNano())
			color := rand.Intn(16777215)

			embed := &discordgo.MessageEmbed{Color: color, Title: "**" + user.Username + "**" + "\n" + "||" + user.ID + "||", Image: &discordgo.MessageEmbedImage{
				URL: user.AvatarURL(""),
			}} //todo

			dm, err := s.UserChannelCreate(i.Interaction.Member.User.ID)
			if err != nil {
				log.Printf("Create PM error : %s", err)
				continue
			}

			go s.ChannelMessageSendEmbed(dm.ID, embed)
		}

		data = discordgo.InteractionResponseData{
			Content: "Команда выполнена",
			Flags:   discordgo.MessageFlagsEphemeral,
		}
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &data,
		})
	}
}
