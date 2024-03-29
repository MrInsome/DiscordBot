// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.23.0
// source: internal/services/discord/api/grpc/proto/discord.proto

package gen

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	BotService_RegisterCommand_FullMethodName = "/discordgo.BotService/RegisterCommand"
	BotService_DeleteCommand_FullMethodName   = "/discordgo.BotService/DeleteCommand"
)

// BotServiceClient is the client API for BotService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BotServiceClient interface {
	RegisterCommand(ctx context.Context, in *RegisterCommandRequest, opts ...grpc.CallOption) (*RegisterCommandResponse, error)
	DeleteCommand(ctx context.Context, in *DeleteCommandRequest, opts ...grpc.CallOption) (*DeleteCommandResponse, error)
}

type botServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBotServiceClient(cc grpc.ClientConnInterface) BotServiceClient {
	return &botServiceClient{cc}
}

func (c *botServiceClient) RegisterCommand(ctx context.Context, in *RegisterCommandRequest, opts ...grpc.CallOption) (*RegisterCommandResponse, error) {
	out := new(RegisterCommandResponse)
	err := c.cc.Invoke(ctx, BotService_RegisterCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *botServiceClient) DeleteCommand(ctx context.Context, in *DeleteCommandRequest, opts ...grpc.CallOption) (*DeleteCommandResponse, error) {
	out := new(DeleteCommandResponse)
	err := c.cc.Invoke(ctx, BotService_DeleteCommand_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BotServiceServer is the server API for BotService service.
// All implementations must embed UnimplementedBotServiceServer
// for forward compatibility
type BotServiceServer interface {
	RegisterCommand(context.Context, *RegisterCommandRequest) (*RegisterCommandResponse, error)
	DeleteCommand(context.Context, *DeleteCommandRequest) (*DeleteCommandResponse, error)
	mustEmbedUnimplementedBotServiceServer()
}

// UnimplementedBotServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBotServiceServer struct {
}

func (UnimplementedBotServiceServer) RegisterCommand(context.Context, *RegisterCommandRequest) (*RegisterCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterCommand not implemented")
}
func (UnimplementedBotServiceServer) DeleteCommand(context.Context, *DeleteCommandRequest) (*DeleteCommandResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteCommand not implemented")
}
func (UnimplementedBotServiceServer) mustEmbedUnimplementedBotServiceServer() {}

// UnsafeBotServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BotServiceServer will
// result in compilation errors.
type UnsafeBotServiceServer interface {
	mustEmbedUnimplementedBotServiceServer()
}

func RegisterBotServiceServer(s grpc.ServiceRegistrar, srv BotServiceServer) {
	s.RegisterService(&BotService_ServiceDesc, srv)
}

func _BotService_RegisterCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).RegisterCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotService_RegisterCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).RegisterCommand(ctx, req.(*RegisterCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BotService_DeleteCommand_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteCommandRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BotServiceServer).DeleteCommand(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: BotService_DeleteCommand_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BotServiceServer).DeleteCommand(ctx, req.(*DeleteCommandRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BotService_ServiceDesc is the grpc.ServiceDesc for BotService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BotService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "discordgo.BotService",
	HandlerType: (*BotServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterCommand",
			Handler:    _BotService_RegisterCommand_Handler,
		},
		{
			MethodName: "DeleteCommand",
			Handler:    _BotService_DeleteCommand_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "internal/services/discord/api/grpc/proto/discord.proto",
}
