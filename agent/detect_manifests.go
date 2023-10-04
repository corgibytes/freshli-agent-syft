package agent

import (
	"freshli-agent-syft/agent/internal"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (c *agentServer) DetectManifests(in *internal.ProjectLocation, server internal.Agent_DetectManifestsServer) error {
	log.Printf("Will detect manifests in %v", in.GetPath())

	return status.Errorf(codes.Unavailable, "Method DetectManifests currently in progress of being implemented")
}
