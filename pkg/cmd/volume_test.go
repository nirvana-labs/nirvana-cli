// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestComputeVolumesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:volumes", "create",
			"--name", "my-data-volume",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--region", "us-sva-2",
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
			"region: us-sva-2\n" +
			"size: 100\n" +
			"type: abs\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n" +
			"vm_id: 123e4567-e89b-12d3-a456-426614174000\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"compute:volumes", "create",
		)
	})
}

func TestComputeVolumesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:volumes", "update",
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
			"compute:volumes", "update",
			"--volume-id", "volume_id",
		)
	})
}

func TestComputeVolumesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:volumes", "list",
			"--max-items", "10",
			"--project-id", "project_id",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestComputeVolumesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:volumes", "delete",
			"--volume-id", "volume_id",
		)
	})
}

func TestComputeVolumesAttach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:volumes", "attach",
			"--volume-id", "volume_id",
			"--vm-id", "123e4567-e89b-12d3-a456-426614174000",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("vm_id: 123e4567-e89b-12d3-a456-426614174000")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"compute:volumes", "attach",
			"--volume-id", "volume_id",
		)
	})
}

func TestComputeVolumesDetach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:volumes", "detach",
			"--volume-id", "volume_id",
		)
	})
}

func TestComputeVolumesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:volumes", "get",
			"--volume-id", "volume_id",
		)
	})
}
