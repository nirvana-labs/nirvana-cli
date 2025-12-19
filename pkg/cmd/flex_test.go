// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestRPCNodesFlexCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex", "create",
		"--blockchain", "ethereum",
		"--name", "my-ethereum-node",
		"--network", "mainnet",
		"--project-id", "123e4567-e89b-12d3-a456-426614174000",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestRPCNodesFlexUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex", "update",
		"--node-id", "node_id",
		"--name", "my-ethereum-node",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestRPCNodesFlexList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex", "list",
		"--cursor", "cursor",
		"--limit", "10",
		"--project-id", "project_id",
	)
}

func TestRPCNodesFlexDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex", "delete",
		"--node-id", "node_id",
	)
}

func TestRPCNodesFlexGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex", "get",
		"--node-id", "node_id",
	)
}
