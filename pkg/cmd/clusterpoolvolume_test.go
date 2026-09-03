// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNKSClustersPoolsVolumesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:pools:volumes", "list",
			"--max-items", "10",
			"--cluster-id", "cluster_id",
			"--pool-id", "pool_id",
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
