package cmd

import (
	"fmt"
	"freshli-agent-syft/agent"
	"github.com/spf13/cobra"
	"log"
	"strconv"
)

func init() {
	rootCmd.AddCommand(
		&cobra.Command{
			Use:   "start-server [port]",
			Short: "Starts a gRPC server running the Freshli Agent service",
			Args: func(cmd *cobra.Command, args []string) error {
				if len(args) == 0 {
					// in the case that no arguments are given we want to show the help text and return with a non-0 return code.
					if err := cmd.Help(); err != nil {
						return fmt.Errorf("unable to display help: %w", err)
					}
					return fmt.Errorf("server port is required")
				}
				return cobra.ExactArgs(1)(cmd, args)
			},
			Run: func(cmd *cobra.Command, args []string) {
				port, err := strconv.Atoi(args[0])
				if err != nil {
					if err := cmd.Help(); err != nil {
						log.Fatalf("unable to display help: %w", err)
					}
					log.Fatalf("could not parse port: %s", err)
				} else {
					log.Printf("Starting agent service on port: %d", port)
					agent.StartAgentService(port)
				}
			},
		})
}
