package sub

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func GetVoiceChannelID(userID string, s *discordgo.Session, i *discordgo.InteractionCreate) (string, error) {
	channel, err := s.State.Channel(i.ChannelID)
	if err != nil {
		return "", fmt.Errorf("failed to fetch channel state: %s", err)
	}

	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		return "", fmt.Errorf("failed to fetch guild state: %s", err)
	}

	for _, vs := range guild.VoiceStates {
		if vs.ChannelID != "" {
			if vs.UserID == userID {
				return vs.ChannelID, nil
			}
		}
	}
	return "", fmt.Errorf("user is not in a voice channel")
}
