package main

import (
	"Discord_bot/internal/config"
	"Discord_bot/internal/services"
	"Discord_bot/internal/services/discord"
	"fmt"
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
		log.Fatalf("Ошибка инициализации файла конфигурации : %s", err.Error())
		return
	}

	newServices, err := services.NewServices(conf)
	if err != nil {
		log.Fatalf("Ошибка при создании сервисов: %s", err.Error())
	}

	cord := discord.NewCordSession(newServices)

	err = cord.InitHandlers()
	if err != nil {
		log.Fatalf("Ошибка при инициализации хендлеров: ", err.Error())
		return
	}

	err = cord.InitCommands()
	if err != nil {
		log.Fatalf("Ошибка при инициализации команд: ", err.Error())
		return
	}

	fmt.Println("Бот запущен. Нажмите CTRL-C для выхода.")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
}
