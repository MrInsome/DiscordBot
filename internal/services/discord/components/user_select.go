package components

import "github.com/bwmarrin/discordgo"

func (c *Components) UserSelect(s *discordgo.Session, i *discordgo.InteractionCreate) {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Вы выбрали пользователя <@" + i.MessageComponentData().Values[0] + ">\n" +
				"Теперь нужно выбрать количество кругов",
			Components: []discordgo.MessageComponent{
				discordgo.ActionsRow{
					Components: []discordgo.MessageComponent{
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								ID: "536499965594435585",
							},
							Label:    "Встряхнуть",
							Style:    discordgo.SecondaryButton,
							Disabled: false,
							CustomID: "wakeup.!1.!" + i.MessageComponentData().Values[0],
						},
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								ID: "893613602387423306",
							},
							Label:    "Разбудить",
							Style:    discordgo.SuccessButton,
							Disabled: false,
							CustomID: "wakeup.!2.!" + i.MessageComponentData().Values[0],
						},
						discordgo.Button{
							Emoji: discordgo.ComponentEmoji{
								ID: "536499966190026762",
							},
							Label:    "Уничтожить",
							Style:    discordgo.DangerButton,
							Disabled: false,
							CustomID: "wakeup.!4.!" + i.MessageComponentData().Values[0],
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
