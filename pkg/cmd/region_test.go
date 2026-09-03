// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestRegionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"regions", "list",
			"--max-items", "10",
			"--availability", "live",
			"--compute-vms=true",
			"--cursor", "cursor",
			"--limit", "10",
			"--networking-connect=true",
			"--networking-vpcs=true",
			"--nks-autoscaling=true",
			"--nks-clusters=true",
			"--sort", "sort",
			"--storage-abs=true",
			"--storage-local-nvme=true",
		)
	})
}

func TestRegionsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"regions", "get",
			"--name", "us-sva-2",
		)
	})
}
