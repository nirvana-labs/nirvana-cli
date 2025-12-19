// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestComputeVolumesAvailabilityCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes:availability", "create",
		"--name", "my-data-volume",
		"--size", "100",
		"--type", "nvme",
		"--vm-id", "vm_id",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestComputeVolumesAvailabilityUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:volumes:availability", "update",
		"--volume-id", "volume_id",
		"--name", "my-data-volume",
		"--size", "100",
		"--tag", "production",
		"--tag", "ethereum",
	)
}
