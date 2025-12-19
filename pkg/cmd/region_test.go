// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestRegionsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"regions", "list",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestRegionsGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"regions", "get",
		"--name", "us-wdc-1",
	)
}
