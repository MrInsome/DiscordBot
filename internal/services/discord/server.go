package discord

import (
	"Discord_bot/internal/services"
)

type CordSession struct {
	*services.Services
}

func NewCordSession(services *services.Services) *CordSession {
	return &CordSession{Services: services}
}

func (cord *CordSession) InitHandlers() error {
	//cord.dg.AddHandler(Ready)

	cord.AddHandler(VoiceState)
	cord.AddHandler(VoiceUsersFunc)

	cord.AddHandler(CommandHello)

	err := cord.Open()
	if err != nil {
		return err
	}
	return nil
}
