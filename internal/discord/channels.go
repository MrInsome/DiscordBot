package discord

import (
	"fmt"
)

func (cord *CordSession) GetBotChannels() {
	channels, err := cord.dg.GuildChannels("1065951599891071046")
	if err != nil {
		fmt.Println("Ошибка при получении списка каналов:", err)
		return
	}

	for _, channel := range channels {
		fmt.Printf("%s - %s\n", channel.Name, channel.ID)
	}
}
