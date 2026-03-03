// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestRPCNodesFlexCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex", "create",
		"--api-key", "string",
		"--blockchain", "ethereum",
		"--name", "my-ethereum-node",
		"--network", "mainnet",
		"--project-id", "123e4567-e89b-12d3-a456-426614174000",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestRPCNodesFlexUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex", "update",
		"--api-key", "string",
		"--node-id", "node_id",
		"--name", "my-ethereum-node",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestRPCNodesFlexList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex", "list",
		"--api-key", "string",
		"--project-id", "project_id",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestRPCNodesFlexDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex", "delete",
		"--api-key", "string",
		"--node-id", "node_id",
	)
}

func TestRPCNodesFlexGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex", "get",
		"--api-key", "string",
		"--node-id", "node_id",
	)
}
