package server

import (
	"context"
	"time"

	"github.com/isometry/go-grpc-echo/rpc"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type grpcServer struct {
	rpc.UnimplementedEchoServer
}

func (s *grpcServer) Request(ctx context.Context, req *rpc.EchoRequest) (*rpc.EchoResponse, error) {
	response := rpc.EchoResponse{
		Message:   req.Message,
		Timestamp: timestamppb.New(time.Now()),
	}
	return &response, nil
}

func NewGRPCServer() *grpc.Server {
	gsrv := grpc.NewServer()
	srv := grpcServer{}
	rpc.RegisterEchoServer(gsrv, &srv)
	return gsrv
}
