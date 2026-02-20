// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestComputeVMsVolumesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms:volumes", "list",
		"--vm-id", "vm_id",
		"--cursor", "cursor",
		"--limit", "10",
	)
}
