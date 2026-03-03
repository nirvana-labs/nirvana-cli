// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestComputeVolumesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "create",
		"--api-key", "string",
		"--name", "my-data-volume",
		"--project-id", "123e4567-e89b-12d3-a456-426614174000",
		"--region", "us-wdc-1",
		"--size", "100",
		"--type", "abs",
		"--tag", "production",
		"--tag", "ethereum",
		"--vm-id", "123e4567-e89b-12d3-a456-426614174000",
	)
}

func TestComputeVolumesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "update",
		"--api-key", "string",
		"--volume-id", "volume_id",
		"--name", "my-data-volume",
		"--size", "100",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestComputeVolumesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "list",
		"--api-key", "string",
		"--project-id", "project_id",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestComputeVolumesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "delete",
		"--api-key", "string",
		"--volume-id", "volume_id",
	)
}

func TestComputeVolumesAttach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "attach",
		"--api-key", "string",
		"--volume-id", "volume_id",
		"--vm-id", "123e4567-e89b-12d3-a456-426614174000",
	)
}

func TestComputeVolumesDetach(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "detach",
		"--api-key", "string",
		"--volume-id", "volume_id",
	)
}

func TestComputeVolumesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "get",
		"--api-key", "string",
		"--volume-id", "volume_id",
	)
}
