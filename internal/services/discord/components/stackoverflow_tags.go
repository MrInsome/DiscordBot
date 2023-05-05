package components

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
	"time"
)

func (c *Components) StackTag(s *discordgo.Session, i *discordgo.InteractionCreate) {
	data := i.MessageComponentData()

	const stackoverflowFormat = `https://stackoverflow.com/questions/tagged/%s`

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Here is your stackoverflow URL: " + fmt.Sprintf(stackoverflowFormat, strings.Join(data.Values, "+")),
			Flags:   discordgo.MessageFlagsEphemeral,
		},
	})
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Second) // Doing that so user won't see instant response.
	_, err = s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: "But wait, there is more! You can also auto populate the select menu. Try executing `/selects auto-populated`.",
		Flags:   discordgo.MessageFlagsEphemeral,
	})
	if err != nil {
		panic(err)
	}
}
