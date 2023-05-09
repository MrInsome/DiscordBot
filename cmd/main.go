package main

import (
	"Discord_bot/internal/config"
	"Discord_bot/internal/services"
	"Discord_bot/internal/services/discord"
	"Discord_bot/internal/services/discord/api/grpc"
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

	go grpc.Serve(cord)

	err = cord.InitHandlers()
	if err != nil {
		log.Fatalf("Handlers init error: ", err.Error())
		return
	}

	servers, err := cord.GetBotServers()
	log.Printf("Bot was started on this servers: \n")
	if err != nil {
		log.Printf("Get servers error : %s \n", err)
	}

	for _, guild := range servers {
		log.Printf("name: %s - id: %s\n", guild.Name, guild.ID)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)
	<-quit
	log.Println("Shutdown...")
}
