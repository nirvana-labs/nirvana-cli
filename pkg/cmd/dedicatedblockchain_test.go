// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestRPCNodesDedicatedBlockchainsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "rpc-nodes:dedicated:blockchains", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}
