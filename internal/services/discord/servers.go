package discord

import (
	"fmt"
)

func (cord *CordSession) GetBotServers() {
	guilds, err := cord.UserGuilds(10, "", "")
	if err != nil {
		fmt.Println("Ошибка при получении списка серверов:", err)
		return
	}

	for _, guild := range guilds {
		fmt.Printf("%s - %s\n", guild.Name, guild.ID)
	}
}
