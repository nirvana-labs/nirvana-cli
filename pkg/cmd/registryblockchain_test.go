// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorRegistryBlockchainsList(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:registry:blockchains", "list",
		"--id", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--symbol", "ethereum_mainnet",
	)
}
