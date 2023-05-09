package contracts

import "github.com/bwmarrin/discordgo"

type ErrorContract interface {
	ErrorMessage(err error, i *discordgo.InteractionCreate)
	DoSomething(arg string)
}

type CommandsContract interface {
	RegisterCommands() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	TestButtons(s *discordgo.Session, i *discordgo.InteractionCreate)
	Selects(s *discordgo.Session, i *discordgo.InteractionCreate)
	UserVoice(s *discordgo.Session, i *discordgo.InteractionCreate)
}

type ComponentsContract interface {
	RegisterComponents() map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate)
	UserSelect(s *discordgo.Session, i *discordgo.InteractionCreate)
	FdNo(s *discordgo.Session, i *discordgo.InteractionCreate)
	FdYes(s *discordgo.Session, i *discordgo.InteractionCreate)
	Select(s *discordgo.Session, i *discordgo.InteractionCreate)
	StackTag(s *discordgo.Session, i *discordgo.InteractionCreate)
}
