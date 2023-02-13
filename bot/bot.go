package bot

import (
	"Discord_bot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	dg, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("Ошибка при создании Дискорд сессии: %v", err)
	}

	dg.AddHandler(messageCreate)

	err = dg.Open()
	if err != nil {
		log.Fatalf("Ошибка при создании соединения: %v", err)
	}

	fmt.Println("Бот запущен.  Нажмите 'CTRL-C' для выхода.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}
