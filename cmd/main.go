package main

import (
	"Discord_bot/internal/config"
	"Discord_bot/internal/services"
	"Discord_bot/internal/services/discord"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func init() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetLevel(logrus.DebugLevel)
}

func main() {
	conf, err := config.Init("config")
	if err != nil {
		log.Fatalf("Config file error : %s", err.Error())
		return
	}

	newServices, err := services.NewServices(conf)
	if err != nil {
		log.Fatalf("Service create error: %s", err.Error())
	}

	cord := discord.NewCordSession(newServices)

	err = cord.InitHandlers()
	if err != nil {
		log.Fatalf("Handlers init error: ", err.Error())
		return
	}

	err = cord.InitCommands()
	if err != nil {
		log.Fatalf("Cannot create slash commands: %v ", err.Error())
		return
	}

	servers, err := cord.GetBotServers()
	log.Printf("Bot was started on this servers: \n")
	if err != nil {
		log.Printf("Get servers error : %s \n", err)
	}

	for _, guild := range servers {
		log.Printf("name: %s - id: %s\n", guild.Name, guild.ID)
		a, _ := cord.GetBotVoiceChannels(guild.ID)
		log.Printf("%s\n", a[0])
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit
	log.Println("Shutdown...")
}
