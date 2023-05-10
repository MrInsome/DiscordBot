package commands

import (
	"Discord_bot/internal/services/discord"
	"Discord_bot/internal/services/discord/api/grpc/gen"
	"context"
	"log"
)

type botService struct {
	*gen.UnimplementedBotServiceServer
	*discord.CordSession
}

func NewBotService(session *discord.CordSession) *botService {
	return &botService{CordSession: session}
}

func (s *botService) RegisterCommand(ctx context.Context, req *gen.RegisterCommandRequest) (*gen.RegisterCommandResponse, error) {
	err := s.InitCom(req.Name, req.Description, req.GuildID, req.Permissions)
	if err != nil {
		log.Printf("Cannot create slash command: %v ", err.Error())
		return &gen.RegisterCommandResponse{Success: false}, err
	}
	log.Println("Command created")
	return &gen.RegisterCommandResponse{Success: true}, nil
}

func (s *botService) DeleteCommand(ctx context.Context, req *gen.DeleteCommandRequest) (*gen.DeleteCommandResponse, error) {
	err := s.DeleteCom(req.Name, req.GuildID)
	if err != nil {
		log.Printf("Cannot delete slash command: %v ", err.Error())
		return &gen.DeleteCommandResponse{Success: false}, err
	}
	log.Println("Command deleted")
	return &gen.DeleteCommandResponse{Success: true}, nil
}
