// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNKSClustersPoolsNodesVolumesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:pools:nodes:volumes", "list",
			"--max-items", "10",
			"--cluster-id", "cluster_id",
			"--pool-id", "pool_id",
			"--node-id", "node_id",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestNKSClustersPoolsNodesVolumesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:pools:nodes:volumes", "get",
			"--cluster-id", "cluster_id",
			"--pool-id", "pool_id",
			"--node-id", "node_id",
			"--volume-id", "volume_id",
		)
	})
}
