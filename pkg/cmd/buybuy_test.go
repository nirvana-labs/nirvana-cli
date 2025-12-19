// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorBuyBuyCreate(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:buy:buy", "create",
		"--blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--from", "0x6b175474e89094c44da98b954eedeac495271d0f",
		"--receive-amount", "10.0000000000000024",
		"--receive-asset", "asset_01jbz9qc00f8wr64hfe459gb7y",
		"--spend-asset", "asset_01jbz9qc00f8wr64hfe459gb7y",
		"--venue", "venue_01jbz9qc18evw86sg8m0sj9jg5",
		"--slippage", "10.0000000000000024",
	)
}
