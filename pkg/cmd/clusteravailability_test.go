// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNKSClustersAvailabilityCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:availability", "create",
			"--name", "my-cluster",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--region", "us-sva-2",
			"--vpc-id", "123e4567-e89b-12d3-a456-426614174000",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-cluster\n" +
			"project_id: 123e4567-e89b-12d3-a456-426614174000\n" +
			"region: us-sva-2\n" +
			"vpc_id: 123e4567-e89b-12d3-a456-426614174000\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"nks:clusters:availability", "create",
		)
	})
}

func TestNKSClustersAvailabilityUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:availability", "update",
			"--cluster-id", "cluster_id",
			"--name", "my-cluster",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-cluster\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"nks:clusters:availability", "update",
			"--cluster-id", "cluster_id",
		)
	})
}
