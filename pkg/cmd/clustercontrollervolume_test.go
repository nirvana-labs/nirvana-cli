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
			"--kind", "boot",
			"--limit", "10",
			"--name", "name",
			"--size-gb-max", "0",
			"--size-gb-min", "0",
			"--sort", "sort",
			"--status", "ready",
			"--type", "abs",
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
