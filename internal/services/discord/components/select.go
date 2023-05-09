package components

import (
	"github.com/bwmarrin/discordgo"
	"log"
	"time"
)

func (c *Components) Select(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var response *discordgo.InteractionResponse

	data := i.MessageComponentData()
	log.Println(data.Values[0])
	switch data.Values[0] {
	case "go":
		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "This is the way.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		}
	default:
		response = &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "It is not the way to go.",
				Flags:   discordgo.MessageFlagsEphemeral,
			},
		}
	}
	err := s.InteractionRespond(i.Interaction, response)
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second) // Doing that so user won't see instant response.
	_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: "Anyways, now when you know how to use single select menus, let's see how multi select menus work. " +
			"Try calling `/selects multi` command.",
		Flags: discordgo.MessageFlagsEphemeral,
	})
	if err != nil {
		panic(err)
	}
}
