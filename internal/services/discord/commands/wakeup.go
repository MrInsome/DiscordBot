package commands

import "github.com/bwmarrin/discordgo"

func (c *Commands) WakeupCommand(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse
	response = &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Необходимо выбрать пользователя которого нужно разбудить.",
			Flags:   discordgo.MessageFlagsEphemeral,
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.SelectMenu{
							MenuType:     discordgo.UserSelectMenu,
							CustomID:     "user_select",
							Placeholder:  "Выбери пользователя",
							MaxValues:    1,
							ChannelTypes: []discordgo.ChannelType{discordgo.ChannelTypeGuildVoice},
						},
					},
				},
			},
		},
	}
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		panic(err)
	}
}
