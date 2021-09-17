# SkySQL API CLI Client

This project contains a command line client for the SkySQL API.

## Usage

The `skysqlcli` may be configured via command line flags, environment variables, or config file (in that order of precedence).

To view the commands that are available, use...

```
skysqlcli --help
```

### Configuration

Environment variables read by this cli client are all prefixed with `SKYSQL_`, and use the same flags as the command line with any hyphens replaced with underscores and using uppercase. For example, to provide the `--api-key` flag via an environment variable, export the value using `SKYSQL_API_KEY`. Similarly, config file entries follow the same pattern as the command line flags with any hyphens replaced with underscores and using lowercase.

### Completion

Completion scripts may be generated for you preferred shell. Instructions are provided for each shell in the respective `--help` output (e.g. `skysqlcli completion bash --help`). Note that autocompletion on the cli will only work when an api-key is provided.

## Contributing

A devcontainer is provided to assist with setting up a developer environment, and is expected to be used with VSCode. Once in the container, run `make deps` to install the dependencies for the project. A `launch.json` is provided as an example for using the debugger.
