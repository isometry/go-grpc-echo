package main

import (
	"log"
	"net"

	"github.com/isometry/go-grpc-echo/internal/server"
	"google.golang.org/grpc/reflection"
)

func main() {
	port := ":8080"

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("Listening on %s", port)
	srv := server.NewGRPCServer()
	reflection.Register(srv)

	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
