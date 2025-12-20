// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestRPCNodesDedicatedList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:dedicated", "list",
		"--cursor", "cursor",
		"--limit", "10",
		"--project-id", "project_id",
	)
}

func TestRPCNodesDedicatedGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:dedicated", "get",
		"--node-id", "node_id",
	)
}
