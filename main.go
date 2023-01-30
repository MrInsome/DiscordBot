package main

import (
	"Discord_bot/bot"
	"Discord_bot/config"
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
