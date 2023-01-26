package bot

import (
	"Discord_bot/config"
	"fmt"
	"github.com/bwmarrin/discordgo"
)

var BotID string
var goBot *discordgo.Session

func Start() {
	fmt.Println("Запуск бота...")
	goBot, err := discordgo.New("Bot " + config.Token)

	if err != nil {
		fmt.Println("Ошибка: ", err.Error())
		return
	}
	user, err := goBot.User("@me")

	if err != nil {
		fmt.Println("Ошибка юзера: ", err.Error())
	}

	BotID = user.ID

	goBot.AddHandler(messageHandler)

	goBot.Identify.Intents = discordgo.IntentsGuildMessages

	err = goBot.Open()

	if err != nil {
		fmt.Println("ошибка: ", err.Error())
		return
	}
	fmt.Println("Бот успешно запущен")
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == BotID {
		return
	}

	if m.Content == "ping" {
		_, _ = s.ChannelMessageSend(m.ChannelID, "@"+m.Author.Username+" Pong!")
	}
}
