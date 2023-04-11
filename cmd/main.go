package main

import (
	"Discord_bot/pkg/bot"
	"Discord_bot/pkg/config"
	"fmt"
)

func main() {
	err := config.ReadConfig()

	if err != nil {
		fmt.Println(err.Error)
		return
	}
	bot.Start()

}
