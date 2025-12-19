// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorBorrowAccountsList(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:borrow:accounts", "list",
		"--account", "0x6b175474e89094c44da98b954eedeac495271d0f",
		"--blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--venue", "venue_01jbz9qc18evw86sg8m0sj9jg5",
		"--at", "2021-01-01T12:00:00Z",
		"--quote-asset-symbol", "eth",
	)
}

func TestVektorBorrowAccountsListHistorical(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:borrow:accounts", "list-historical",
		"--account", "0x6b175474e89094c44da98b954eedeac495271d0f",
		"--blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--from", "2021-01-01T12:00:00Z",
		"--to", "2021-01-01T12:00:00Z",
		"--venue", "venue_01jbz9qc18evw86sg8m0sj9jg5",
		"--quote-asset-symbol", "eth",
	)
}
