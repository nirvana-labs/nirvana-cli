// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestComputeVolumesAvailabilityCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:volumes:availability", "create",
			"--name", "my-data-volume",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--region", "us-wdc-1",
			"--size", "100",
			"--type", "abs",
			"--tag", "production",
			"--tag", "ethereum",
			"--vm-id", "123e4567-e89b-12d3-a456-426614174000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-data-volume\n" +
			"project_id: 123e4567-e89b-12d3-a456-426614174000\n" +
			"region: us-wdc-1\n" +
			"size: 100\n" +
			"type: abs\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n" +
			"vm_id: 123e4567-e89b-12d3-a456-426614174000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"compute:volumes:availability", "create",
		)
	})
}

func TestComputeVolumesAvailabilityUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:volumes:availability", "update",
			"--volume-id", "volume_id",
			"--name", "my-data-volume",
			"--size", "100",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-data-volume\n" +
			"size: 100\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"compute:volumes:availability", "update",
			"--volume-id", "volume_id",
		)
	})
}
