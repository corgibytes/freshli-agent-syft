# Freshli Agent: Java

This application is used by the [`freshli` CLI](https://github.com/corgibytes/freshli-cli) to detect and process manifest files with the Syft processor.

This iteration of the agent simply wraps the [Syft cli](https://github.com/anchore/syft) as described in the Syft [README](https://github.com/anchore/syft#readme).

Invoked via: `go run cmd/agent/main.go`

