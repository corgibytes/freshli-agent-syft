package agent

import (
	"fmt"
	"freshli-agent-syft/agent/internal"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

var grpcServer = grpc.NewServer()

type agentServer struct {
	internal.UnimplementedAgentServer
}

func StartAgentService(port int) {

	log.Printf("gRPC server listening on %d", port)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	healthcheck := health.NewServer()
	healthgrpc.RegisterHealthServer(grpcServer, healthcheck)
	internal.RegisterAgentServer(grpcServer, &agentServer{})
	healthcheck.SetServingStatus(internal.Agent_ServiceDesc.ServiceName, healthpb.HealthCheckResponse_SERVING)

	// TODO should this be only in dev/debug mode?
	reflection.Register(grpcServer)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
