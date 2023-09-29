package cmd

import (
	"github.com/anchore/clio"
	"github.com/anchore/syft/cmd/syft/cli"
)

const (
	NotProvided = "[not provided]"
)

const applicationName = "syft"

func init() {
	rootCmd.AddCommand(
		cli.Command(clio.Identification{
			Name:           applicationName,
			Version:        NotProvided,
			BuildDate:      NotProvided,
			GitCommit:      NotProvided,
			GitDescription: NotProvided,
		}),
	)
}
