package components

import "github.com/bwmarrin/discordgo"

func (c *Components) ChanSelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "This is it. You've reached your destination. Your choice was <#" + i.MessageComponentData().Values[0] + ">\n" +
				"If you want to know more, check out the links below",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								Name: "📜",
							},
							Label: "Documentation",
							Style: discordgo.LinkButton,
							URL:   "https://discord.com/developers/docs/interactions/message-components#select-menus",
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

			Flags: discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		panic(err)
	}
}
