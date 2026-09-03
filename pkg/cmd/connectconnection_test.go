// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestNetworkingConnectConnectionsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:connect:connections", "create",
			"--bandwidth-mbps", "50",
			"--cidr", "10.0.0.0/16",
			"--name", "my-connect-connection",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--provider-cidr", "172.16.0.0/16",
			"--region", "us-sva-2",
			"--aws", "{account_id: '523816707215', region: us-east-1}",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(networkingConnectConnectionsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:connect:connections", "create",
			"--bandwidth-mbps", "50",
			"--cidr", "10.0.0.0/16",
			"--name", "my-connect-connection",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--provider-cidr", "172.16.0.0/16",
			"--region", "us-sva-2",
			"--aws.account-id", "523816707215",
			"--aws.region", "us-east-1",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"bandwidth_mbps: 50\n" +
			"cidrs:\n" +
			"  - 10.0.0.0/16\n" +
			"name: my-connect-connection\n" +
			"project_id: 123e4567-e89b-12d3-a456-426614174000\n" +
			"provider_cidrs:\n" +
			"  - 172.16.0.0/16\n" +
			"region: us-sva-2\n" +
			"aws:\n" +
			"  account_id: '523816707215'\n" +
			"  region: us-east-1\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"networking:connect:connections", "create",
		)
	})
}

func TestNetworkingConnectConnectionsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:connect:connections", "update",
			"--connection-id", "connection_id",
			"--name", "my-connect-connection",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-connect-connection\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"networking:connect:connections", "update",
			"--connection-id", "connection_id",
		)
	})
}

func TestNetworkingConnectConnectionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:connect:connections", "list",
			"--max-items", "10",
			"--project-id", "project_id",
			"--bandwidth-mbps", "50",
			"--cursor", "cursor",
			"--limit", "10",
			"--name", "name",
			"--provider", "provider",
			"--provider-region", "provider_region",
			"--region", "region",
			"--sort", "sort",
			"--status", "ready",
			"--tag", "string",
		)
	})
}

func TestNetworkingConnectConnectionsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:connect:connections", "delete",
			"--connection-id", "connection_id",
		)
	})
}

func TestNetworkingConnectConnectionsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:connect:connections", "get",
			"--connection-id", "connection_id",
		)
	})
}
