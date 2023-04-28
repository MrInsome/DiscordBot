package discord

import (
	"github.com/bwmarrin/discordgo"
)

func (cord *CordSession) GetBotAllChannels(guildID string) ([]string, error) {
	var allChannels []string
	channels, err := cord.GuildChannels(guildID)
	if err != nil {
		return nil, err
	}

	for _, channel := range channels {
		allChannels = append(allChannels, channel.ID)
	}
	return allChannels, err
}

func (cord *CordSession) GetBotTextChannels(guildID string) ([]string, error) {
	var textChannels []string
	channels, err := cord.GuildChannels(guildID)
	if err != nil {
		return nil, err
	}

	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildText {
			textChannels = append(textChannels, channel.ID)
		}
	}
	return textChannels, nil
}

func (cord *CordSession) GetBotVoiceChannels(guildID string) ([]string, error) {
	var voiceChannels []string
	channels, err := cord.GuildChannels(guildID)
	if err != nil {
		return nil, err
	}

	for _, channel := range channels {
		if channel.Type == discordgo.ChannelTypeGuildVoice {
			voiceChannels = append(voiceChannels, channel.ID)
		}
	}
	return voiceChannels, nil
}
