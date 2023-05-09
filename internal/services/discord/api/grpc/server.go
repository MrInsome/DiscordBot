package grpc

import (
	"Discord_bot/internal/services/discord"
	"Discord_bot/internal/services/discord/api/grpc/commands"
	"Discord_bot/internal/services/discord/api/grpc/gen"
	"google.golang.org/grpc"
	"log"
	"net"
)

func Serve(s *discord.CordSession) {
	// создаем gRPC-сервер
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()

	// регистрируем сервис BotService
	gen.RegisterBotServiceServer(grpcServer, commands.NewBotService(s))

	// запускаем сервер
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
