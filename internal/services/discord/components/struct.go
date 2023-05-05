package components

import (
	"Discord_bot/internal/contracts"
	"github.com/bwmarrin/discordgo"
)

type Components struct {
	*discordgo.Session
	contracts.ComponentsContract
}

func NewComponents(dg *discordgo.Session) *Components {
	return &Components{Session: dg}
}

func (c *Components) RegisterComponents() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate) {
	var componentsHandlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
		"fd_no":              c.FdNo,
		"fd_yes":             c.FdYes,
		"select":             c.Select,
		"stackoverflow_tags": c.StackTag,
		"channel_select":     c.ChanSelect,
	}
	return componentsHandlers
}
