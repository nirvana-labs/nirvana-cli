// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorLendSetCollateralCreate(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:lend:set-collateral", "create",
		"--blockchain", "blockchain_01jbz9nsy8egar70jg79dkwmaf",
		"--from", "0x6b175474e89094c44da98b954eedeac495271d0f",
		"--market-id", "lend_borrow_market_01h455vb4pex5vsknk084sn02q",
		"--status",
	)
}
