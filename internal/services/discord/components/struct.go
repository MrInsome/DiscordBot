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
		"user_select": c.UserSelect,
		"wakeup":      c.Wakeuper,

		"fd_no":              c.FdNo,     //todo
		"fd_yes":             c.FdYes,    //todo
		"select":             c.Select,   //todo
		"stackoverflow_tags": c.StackTag, //todo
	}
	return componentsHandlers
}
