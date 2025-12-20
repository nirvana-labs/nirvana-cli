// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestComputeVMsOSImagesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms:os-images", "list",
		"--cursor", "cursor",
		"--limit", "10",
	)
}
