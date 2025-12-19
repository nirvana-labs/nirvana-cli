// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNetworkingConnectRoutesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:routes", "list",
		"--cursor", "cursor",
		"--limit", "10",
	)
}
