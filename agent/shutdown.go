package agent

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	"log"
)

func (c *agentServer) Shutdown(ctx context.Context, in *empty.Empty) (*empty.Empty, error) {
	log.Printf("shutting down gRPC server")
	if grpcServer != nil {
		go grpcServer.GracefulStop()
	}
	return new(empty.Empty), nil
}
