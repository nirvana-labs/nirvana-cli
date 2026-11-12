// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestSandboxClustersCostCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"sandbox-clusters:cost", "create",
			"--instance-type", "n1-standard-8",
			"--name", "my-sandboxes",
			"--node-count", "3",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--region", "us-sva-2",
			"--tag", "production",
			"--tag", "agents",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instance_type: n1-standard-8\n" +
			"name: my-sandboxes\n" +
			"node_count: 3\n" +
			"project_id: 123e4567-e89b-12d3-a456-426614174000\n" +
			"region: us-sva-2\n" +
			"tags:\n" +
			"  - production\n" +
			"  - agents\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"sandbox-clusters:cost", "create",
		)
	})
}
