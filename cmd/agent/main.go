package main

import (
	_ "modernc.org/sqlite"

	"github.com/anchore/clio"
	"github.com/anchore/syft/cmd/syft/cli"
)

const (
	NotProvided = "[not provided]"
)

const applicationName = "freshli-agent-syft"

func main() {
	app := cli.Application(
		clio.Identification{
			Name:           applicationName,
			Version:        NotProvided,
			BuildDate:      NotProvided,
			GitCommit:      NotProvided,
			GitDescription: NotProvided,
		},
	)

	app.Run()
}
