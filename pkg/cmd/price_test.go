// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorPricesList(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:prices", "list",
		"--asset-symbol", "eth",
		"--quote-asset-symbol", "eth",
	)
}

func TestVektorPricesListHistorical(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:prices", "list-historical",
		"--asset-symbol", "eth",
		"--quote-asset-symbol", "eth",
	)
}
