// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestComputeVolumesCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "create",
		"--name", "my-data-volume",
		"--size", "100",
		"--type", "nvme",
		"--vm-id", "vm_id",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestComputeVolumesUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "update",
		"--volume-id", "volume_id",
		"--name", "my-data-volume",
		"--size", "100",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestComputeVolumesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "list",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestComputeVolumesDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "delete",
		"--volume-id", "volume_id",
	)
}

func TestComputeVolumesAttach(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "attach",
		"--volume-id", "volume_id",
		"--vm-id", "123e4567-e89b-12d3-a456-426614174000",
	)
}

func TestComputeVolumesDetach(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "detach",
		"--volume-id", "volume_id",
	)
}

func TestComputeVolumesGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes", "get",
		"--volume-id", "volume_id",
	)
}
