package discord

import (
	"github.com/bwmarrin/discordgo"
)

func (cord *CordSession) GetBotServers() ([]*discordgo.UserGuild, error) {
	guilds, err := cord.UserGuilds(10, "", "")
	if err != nil {
		return nil, err
	}
	return guilds, err
}
