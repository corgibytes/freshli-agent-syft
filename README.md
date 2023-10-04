# Freshli Agent: Java

This application is used by the [`freshli` CLI](https://github.com/corgibytes/freshli-cli) to detect and process manifest files with the Syft processor.

## Running

There are two commands in this agent: `syft` and `start-server`.

### `syft`

This command simply wraps the [Syft cli](https://github.com/anchore/syft) as described in the Syft [README](https://github.com/anchore/syft#readme).

Simply run: `go run main.go syft`

### `start-server`

This command runs the gRPC server on the specified port.

`go run main.go start-server 9093`
