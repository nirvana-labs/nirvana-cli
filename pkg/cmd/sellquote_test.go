// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorSellQuotesList(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:sell:quotes", "list",
		"--blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--receive-asset", "asset_01jbz9qc00f8wr64hfe459gb7y",
		"--spend-amount", "10.0000000000000024",
		"--spend-asset", "asset_01jbz9qc00f8wr64hfe459gb7y",
		"--venue", "venue_01jbz9qc18evw86sg8m0sj9jg5",
		"--quote-asset-symbol", "eth",
	)
}
