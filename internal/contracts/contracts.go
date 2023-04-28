package contracts

import "github.com/bwmarrin/discordgo"

type ErrorContract interface {
	TestErr(err error, i *discordgo.InteractionCreate)
}
