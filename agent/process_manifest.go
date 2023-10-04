package agent

import (
	"context"
	"freshli-agent-syft/agent/internal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (c *agentServer) ProcessManifest(ctx context.Context, in *internal.ProcessingRequest) (*internal.BomLocation, error) {
	log.Printf("Will process manifest %v for %v", in.GetManifest().GetPath(), in.GetMoment())

	return nil, status.Errorf(codes.Unavailable, "Method ProcessManifest currently in progress of being implemented")
}
