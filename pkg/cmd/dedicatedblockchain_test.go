// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestRPCNodesDedicatedBlockchainsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"rpc-nodes:dedicated:blockchains", "list",
		"--cursor", "cursor",
		"--limit", "10",
	)
}
