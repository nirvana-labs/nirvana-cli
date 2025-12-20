// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorLendPositionsList(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:lend:positions", "list",
		"--account", "0x6b175474e89094c44da98b954eedeac495271d0f",
		"--asset", "asset_01jbz9qc00f8wr64hfe459gb7y",
		"--blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--venue", "venue_01jbz9qc18evw86sg8m0sj9jg5",
		"--at", "2021-01-01T12:00:00Z",
		"--quote-asset-symbol", "eth",
		"--market-id", "lend_borrow_market_01h455vb4pex5vsknk084sn02q",
	)
}

func TestVektorLendPositionsListHistorical(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:lend:positions", "list-historical",
		"--account", "0x6b175474e89094c44da98b954eedeac495271d0f",
		"--asset", "asset_01jbz9qc00f8wr64hfe459gb7y",
		"--blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--from", "2021-01-01T12:00:00Z",
		"--to", "2021-01-01T12:00:00Z",
		"--venue", "venue_01jbz9qc18evw86sg8m0sj9jg5",
		"--quote-asset-symbol", "eth",
		"--market-id", "lend_borrow_market_01h455vb4pex5vsknk084sn02q",
	)
}
