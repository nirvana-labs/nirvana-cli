// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestNKSClustersPoolsAvailabilityCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:pools:availability", "create",
			"--cluster-id", "cluster_id",
			"--name", "my-node-pool",
			"--node-config", "{ram_gi: 8, storage_gi: 100, vcpu: 4}",
			"--node-count", "3",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(nksClustersPoolsAvailabilityCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:pools:availability", "create",
			"--cluster-id", "cluster_id",
			"--name", "my-node-pool",
			"--node-config.ram-gi", "8",
			"--node-config.storage-gi", "100",
			"--node-config.vcpu", "4",
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
			"  ram_gi: 8\n" +
			"  storage_gi: 100\n" +
			"  vcpu: 4\n" +
			"node_count: 3\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"nks:clusters:pools:availability", "create",
			"--cluster-id", "cluster_id",
		)
	})
}

func TestNKSClustersPoolsAvailabilityUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:pools:availability", "update",
			"--cluster-id", "cluster_id",
			"--pool-id", "pool_id",
			"--name", "my-node-pool",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-node-pool\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"nks:clusters:pools:availability", "update",
			"--cluster-id", "cluster_id",
			"--pool-id", "pool_id",
		)
	})
}
