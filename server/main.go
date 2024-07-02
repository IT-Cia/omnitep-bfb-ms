package main

import (
	"gsm/config"
	"gsm/server/task"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	log.Println("server starting...")
	lis, err := net.Listen("tcp", config.SERVER_GRPC_ADDRESS)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)

	task.Register(grpcServer)

	log.Println("server listening to ", config.SERVER_GRPC_ADDRESS)
	grpcServer.Serve(lis)
}
