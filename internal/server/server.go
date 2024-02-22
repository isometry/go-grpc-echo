package server

import (
	"context"
	"time"

	api "github.com/isometry/go-grpc-echo/api/v1"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type grpcServer struct {
	api.UnimplementedEchoServer
}

func (s *grpcServer) Echo(ctx context.Context, req *api.EchoRequest) (*api.EchoResponse, error) {
	response := api.EchoResponse{
		Message:   req.Message,
		Timestamp: timestamppb.New(time.Now()),
	}
	return &response, nil
}

func NewGRPCServer() *grpc.Server {
	gsrv := grpc.NewServer()
	srv := grpcServer{}
	api.RegisterEchoServer(gsrv, &srv)
	return gsrv
}
