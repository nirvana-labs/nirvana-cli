// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestRPCNodesDedicatedList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "rpc-nodes:dedicated", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--project-id", "project_id",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestRPCNodesDedicatedGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "rpc-nodes:dedicated", "get",
			"--api-key", "string",
			"--node-id", "node_id",
		)
	})
}
