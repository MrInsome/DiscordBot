package main

import (
	"Discord_bot/internal/config"
	"Discord_bot/internal/discord"
	"fmt"
	"log"
)

func main() {
	conf, err := config.Init("config")
	if err != nil {
		log.Fatalf("Ошибка инициализации файла конфигурации : %s", err.Error())
		return
	}
	dg, err := discord.NewDiscordSession(conf)
	if err != nil {
		log.Fatalf("Ошибка при создании сессии Discord: %s", err.Error())
		return
	}

	//commands.GetBotServers(dg)
	//commands.GetBotChannels(dg)

	err = dg.InitSession()
	if err != nil {
		log.Fatalf("Ошибка при открытии сессии Discord: ", err.Error())
		return
	}
	dg.InitCommands()
	fmt.Println("Бот запущен. Нажмите CTRL-C для выхода.")
	<-make(chan struct{})
}
