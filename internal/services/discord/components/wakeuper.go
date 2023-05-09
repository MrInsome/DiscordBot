package components

import (
	"Discord_bot/internal/services/discord/commands/sub"
	"github.com/bwmarrin/discordgo"
	"strings"
)

func (c *Components) Wakeuper(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := strings.Split(i.MessageComponentData().CustomID, ".!")
	voiceID, err := sub.GetVoiceChannelID(data[2], s, i)
	if err != nil {
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: err.Error(),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Пользователь найден в канале<#" + voiceID + ">.\n" + "Вариант " + data[1] + " в процессе...",
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	go sub.WakeupCycle(data[2], voiceID, data[1], s, i)
	if err != nil {
		err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: err.Error(),
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		})
	}
}
