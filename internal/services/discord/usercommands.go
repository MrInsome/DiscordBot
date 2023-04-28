package discord

import (
	"github.com/bwmarrin/discordgo"
	"log"
)

func Ready(dg *discordgo.Session, event *discordgo.Ready) {
	_, err := dg.ChannelMessageSend("1065951600541179987", "Бот запущен")
	if err != nil {
		log.Printf("Ошибка при отправке сообщения в канал: %s", err)
	}
}

func CommandHello(dg *discordgo.Session, i *discordgo.InteractionCreate) {
	if i.ApplicationCommandData().Name == "hello" {
		dg.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Hello, World!",
			},
		})
	}
}
