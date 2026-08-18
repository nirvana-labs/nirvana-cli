// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestComputeVMsMetricsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:vms:metrics", "list",
			"--vm-id", "vm_id",
			"--aggregation", "mean",
			"--end-time", "'2019-12-27T18:11:19.117Z'",
			"--metric", "compute.nirvanalabs.io/vm/cpu/used_cores",
			"--period", "5m",
			"--start-time", "'2019-12-27T18:11:19.117Z'",
		)
	})
}
