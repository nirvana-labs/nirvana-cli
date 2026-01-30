# Nirvana Labs CLI

The official CLI for the [Nirvana Labs REST API](https://docs.nirvanalabs.io).

<!-- x-release-please-start-version -->

## Installation

### Installing with Homebrew

```sh
brew tap nirvana-labs/tap
brew install nirvana
```

### Installing with Go

To test or install the CLI locally, you need [Go](https://go.dev/doc/install) version 1.22 or later installed.

```sh
go install 'github.com/nirvana-labs/nirvana-cli/cmd/nirvana@latest'
```

Once you have run `go install`, the binary is placed in your Go bin directory:

- **Default location**: `$HOME/go/bin` (or `$GOPATH/bin` if GOPATH is set)
- **Check your path**: Run `go env GOPATH` to see the base directory

If commands aren't found after installation, add the Go bin directory to your PATH:

```sh
# Add to your shell profile (.zshrc, .bashrc, etc.)
export PATH="$PATH:$(go env GOPATH)/bin"
```

<!-- x-release-please-end -->

### Running Locally

After cloning the git repository for this project, you can use the
`scripts/run` script to run the tool locally:

```sh
./scripts/run args...
```

## Usage

The CLI follows a resource-based command structure:

```sh
nirvana [resource] [command] [flags]
```

```sh
nirvana compute:vms create \
  --boot-volume '{size: 100, type: nvme, tags: [production, ethereum]}' \
  --cpu-config '{vcpu: 2}' \
  --memory-config '{size: 2}' \
  --name my-vm \
  --os-image-name ubuntu-noble-2025-10-01 \
  --public-ip-enabled \
  --region us-wdc-1 \
  --ssh-key '{public_key: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2}' \
  --subnet-id 123e4567-e89b-12d3-a456-426614174000 \
  --data-volume '{name: my-data-volume, size: 100, type: nvme, tags: [production, ethereum]}' \
  --tag production \
  --tag ethereum
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--help` - Show command line usage
- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
- `--base-url` - Use a custom API backend URL
- `--format` - Change the output format (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--format-error` - Change the output format for errors (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--transform` - Transform the data output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
- `--transform-error` - Transform the error output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
