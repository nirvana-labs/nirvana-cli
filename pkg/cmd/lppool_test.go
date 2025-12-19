// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorLPPoolsList(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:lp:pools", "list",
		"--asset", "asset_01jbz9qc00f8wr64hfe459gb7y",
		"--blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--venue", "venue_01jbz9qc18evw86sg8m0sj9jg5",
		"--at", "2021-01-01T12:00:00Z",
		"--quote-asset-symbol", "eth",
		"--lp-pool-id", "lp_pool_01h455vb4pex5vsknk084sn02q",
	)
}

func TestVektorLPPoolsListHistorical(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:lp:pools", "list-historical",
		"--asset", "asset_01jbz9qc00f8wr64hfe459gb7y",
		"--blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--from", "2021-01-01T12:00:00Z",
		"--to", "2021-01-01T12:00:00Z",
		"--venue", "venue_01jbz9qc18evw86sg8m0sj9jg5",
		"--quote-asset-symbol", "eth",
		"--lp-pool-id", "lp_pool_01h455vb4pex5vsknk084sn02q",
	)
}
