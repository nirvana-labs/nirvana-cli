// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorRegistryBorrowMarketsList(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:registry:borrow-markets", "list",
		"--id", "lend_borrow_market_01h455vb4pex5vsknk084sn02q",
		"--asset", "asset_01jbz9qc00f8wr64hfe459gb7y",
		"--blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--venue", "venue_01jbz9qc18evw86sg8m0sj9jg5",
	)
}
