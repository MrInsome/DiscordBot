package bot

import (
	"github.com/bwmarrin/discordgo"
)

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	if m.Content == "pong" {
		s.ChannelMessageSend(m.ChannelID, "Ping!")
	}
	if m.Content == "!role" {
		guild, err := s.State.Guild(m.GuildID)
		if err != nil {
			return
		}

		for _, role := range guild.Roles {
			if role.Name == "admin" {
				err := s.GuildMemberRoleAdd(guild.ID, m.Author.ID, role.ID)
				if err != nil {
					return
				}
				return
			}
		}
	}
}
