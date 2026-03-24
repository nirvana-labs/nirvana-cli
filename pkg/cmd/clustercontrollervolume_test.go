// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNKSClustersControllersVolumesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:controllers:volumes", "list",
			"--max-items", "10",
			"--cluster-id", "cluster_id",
			"--controller-id", "controller_id",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestNKSClustersControllersVolumesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:controllers:volumes", "get",
			"--cluster-id", "cluster_id",
			"--controller-id", "controller_id",
			"--volume-id", "volume_id",
		)
	})
}
