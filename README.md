# Nirvana Labs CLI

The official CLI for the [Nirvana Labs REST API](https://docs.nirvanalabs.io).

## Installation

### Installing with Go

<!-- x-release-please-start-version -->

```sh
go install 'github.com/nirvana-labs/nirvana-cli/cmd/nirvana@latest'
```

### Running Locally

<!-- x-release-please-start-version -->

```sh
go run cmd/nirvana/main.go
```

<!-- x-release-please-end -->

## Usage

The CLI follows a resource-based command structure:

```sh
nirvana [resource] [command] [flags]
```

```sh
nirvana compute:vms create \
  --boot-volume '{size: 100, type: nvme}' \
  --cpu-config '{vcpu: 2}' \
  --memory-config '{size: 2}' \
  --name my-vm \
  --os-image-name ubuntu-noble-2025-10-01 \
  --public-ip-enabled \
  --region us-wdc-1 \
  --ssh-key '{public_key: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2}' \
  --subnet-id 123e4567-e89b-12d3-a456-426614174000
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
