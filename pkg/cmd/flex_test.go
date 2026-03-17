// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestRPCNodesFlexCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rpc-nodes:flex", "create",
			"--blockchain", "ethereum",
			"--name", "my-ethereum-node",
			"--network", "mainnet",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"blockchain: ethereum\n" +
			"name: my-ethereum-node\n" +
			"network: mainnet\n" +
			"project_id: 123e4567-e89b-12d3-a456-426614174000\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"rpc-nodes:flex", "create",
		)
	})
}

func TestRPCNodesFlexUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rpc-nodes:flex", "update",
			"--node-id", "node_id",
			"--name", "my-ethereum-node",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-ethereum-node\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"rpc-nodes:flex", "update",
			"--node-id", "node_id",
		)
	})
}

func TestRPCNodesFlexList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rpc-nodes:flex", "list",
			"--max-items", "10",
			"--project-id", "project_id",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestRPCNodesFlexDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rpc-nodes:flex", "delete",
			"--node-id", "node_id",
		)
	})
}

func TestRPCNodesFlexGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"rpc-nodes:flex", "get",
			"--node-id", "node_id",
		)
	})
}
