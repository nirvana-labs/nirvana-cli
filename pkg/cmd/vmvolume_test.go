// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestComputeVMsVolumesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:vms:volumes", "list",
			"--max-items", "10",
			"--vm-id", "vm_id",
			"--cursor", "cursor",
			"--kind", "boot",
			"--limit", "10",
			"--name", "name",
			"--sort", "sort",
			"--status", "pending",
			"--tag", "string",
			"--type", "nvme",
		)
	})
}
