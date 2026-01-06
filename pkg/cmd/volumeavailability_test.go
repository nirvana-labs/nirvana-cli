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
		"--region", "us-wdc-1",
		"--size", "100",
		"--type", "nvme",
		"--tag", "production",
		"--tag", "ethereum",
		"--vm-id", "123e4567-e89b-12d3-a456-426614174000",
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
