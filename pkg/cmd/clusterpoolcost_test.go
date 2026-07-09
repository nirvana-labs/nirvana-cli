// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestNKSClustersPoolsCostCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:pools:cost", "create",
			"--cluster-id", "cluster_id",
			"--name", "my-node-pool",
			"--node-config", "{boot_volume: {size: 100, type: abs}, instance_type: n1-standard-8, labels: [env=prod, team=platform], taints: [dedicated=gpu:NoSchedule]}",
			"--node-count", "3",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(nksClustersPoolsCostCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:pools:cost", "create",
			"--cluster-id", "cluster_id",
			"--name", "my-node-pool",
			"--node-config.boot-volume", "{size: 100, type: abs}",
			"--node-config.instance-type", "n1-standard-8",
			"--node-config.labels", "[env=prod, team=platform]",
			"--node-config.taints", "[dedicated=gpu:NoSchedule]",
			"--node-count", "3",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-node-pool\n" +
			"node_config:\n" +
			"  boot_volume:\n" +
			"    size: 100\n" +
			"    type: abs\n" +
			"  instance_type: n1-standard-8\n" +
			"  labels:\n" +
			"    - env=prod\n" +
			"    - team=platform\n" +
			"  taints:\n" +
			"    - dedicated=gpu:NoSchedule\n" +
			"node_count: 3\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"nks:clusters:pools:cost", "create",
			"--cluster-id", "cluster_id",
		)
	})
}

func TestNKSClustersPoolsCostUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:pools:cost", "update",
			"--cluster-id", "cluster_id",
			"--pool-id", "pool_id",
			"--name", "my-node-pool",
			"--node-count", "5",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-node-pool\n" +
			"node_count: 5\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"nks:clusters:pools:cost", "update",
			"--cluster-id", "cluster_id",
			"--pool-id", "pool_id",
		)
	})
}
