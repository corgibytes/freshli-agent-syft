package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"strconv"
)

func main() {
	port, err := parsePortArgument()
	if err != nil || port < 0 {
		log.Fatalf("port is required: %v", err)
		return
	}
	log.Printf("gRPC server listening on %d", port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(s, healthcheck)

	// TODO should this be only in dev/debug mode?
	reflection.Register(s)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func parsePortArgument() (int, error) {
	if len(os.Args) > 1 {
		return strconv.Atoi(os.Args[1])
	}
	return -1, nil
}
