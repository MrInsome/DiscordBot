package discord

import (
	"Discord_bot/internal/config"
	"github.com/bwmarrin/discordgo"
)

type CordSession struct {
	dg   *discordgo.Session
	conf *config.Configs
}

func NewDiscordSession(conf *config.Configs) (*CordSession, error) {
	dg, err := discordgo.New("Bot " + conf.Token)
	if err != nil {
		return nil, err
	}
	return &CordSession{dg: dg, conf: conf}, nil
}

func (cord *CordSession) InitSession() error {
	cord.dg.AddHandler(VoiceState)
	cord.dg.AddHandler(Ready)
	cord.dg.AddHandler(CommandHello)
	cord.dg.AddHandler(VoiceUsersFunc)

	err := cord.dg.Open()
	if err != nil {
		return err
	}
	return nil
}
