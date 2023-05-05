package components

import "github.com/bwmarrin/discordgo"

func (c *Components) FdNo(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Huh. I see, maybe some of these resources might help you?",
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "📜",
							},
							Label: "Documentation",
							Style: discordgo.LinkButton,
							URL:   "https://discord.com/developers/docs/interactions/message-components#buttons",
						},
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "🔧",
							},
							Label: "Discord developers",
							Style: discordgo.LinkButton,
							URL:   "https://discord.gg/discord-developers",
						},
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "🦫",
							},
							Label: "Discord Gophers",
							Style: discordgo.LinkButton,
							URL:   "https://discord.gg/7RuRrVHyXF",
						},
					},
				},
			},
		},
	})
	if err != nil {
		panic(err)
	}
}
