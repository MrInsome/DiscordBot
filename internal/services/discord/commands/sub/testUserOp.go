package sub

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func GetUserOptions(s *discordgo.Session, i *discordgo.InteractionCreate) []discordgo.SelectMenuOption {
	var options []discordgo.SelectMenuOption

	channel, err := s.State.Channel(i.ChannelID)
	if err != nil {
		//cord.ErrorContract.ErrorMessage(err, i)
		log.Printf("Error from GetGuild : %s", err)
		return nil
	}

	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		//cord.ErrorContract.ErrorMessage(err, i)
		log.Printf("Error from cord GetGuild : %s", err)
		return nil
	}

	usersInVoice := make(map[string][]string)

	for _, vs := range guild.VoiceStates {
		if vs.ChannelID != "" {
			if _, ok := usersInVoice[vs.ChannelID]; !ok {
				usersInVoice[vs.ChannelID] = []string{}
			}
			usersInVoice[vs.ChannelID] = append(usersInVoice[vs.ChannelID], vs.UserID)
			options = append(options, discordgo.SelectMenuOption{
				Label: vs.Member.Nick,
				Value: vs.UserID,
			})
		}
	}

	return options
}
