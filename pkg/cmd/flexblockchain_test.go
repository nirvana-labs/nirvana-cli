// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestRPCNodesFlexBlockchainsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:flex:blockchains", "list",
		"--cursor", "cursor",
		"--limit", "10",
	)
}
