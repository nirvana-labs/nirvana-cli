// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorBridgeQuotesList(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:bridge:quotes", "list",
		"--amount", "10.0000000000000024",
		"--asset", "eth",
		"--from-blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--to-blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--venue", "venue_01jbz9qc18evw86sg8m0sj9jg5",
		"--quote-asset-symbol", "eth",
	)
}
