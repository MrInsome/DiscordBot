package discord

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
)

func Ready(dg *discordgo.Session, event *discordgo.Ready) {
	_, err := dg.ChannelMessageSend("1065951600541179987", "Бот запущен")
	if err != nil {
		fmt.Println("Ошибка при отправке сообщения в канал:", err)
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
