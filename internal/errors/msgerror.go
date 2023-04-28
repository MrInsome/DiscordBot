package errors

import (
	"Discord_bot/internal/contracts"
	"github.com/bwmarrin/discordgo"
)

type DiscordErrors struct {
	*discordgo.Session
	dErr contracts.ErrorContract
}

func NewDiscordErrors() *DiscordErrors {
	return &DiscordErrors{}
}

func (de *DiscordErrors) TestErr(err error, i *discordgo.InteractionCreate) {
	data := discordgo.InteractionResponseData{
		Content: err.Error(),
		Flags:   discordgo.MessageFlagsEphemeral,
	}

	err = de.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &data,
	})
}
