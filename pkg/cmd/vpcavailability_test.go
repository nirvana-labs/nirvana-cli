// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNetworkingVPCsAvailabilityCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:vpcs:availability", "create",
		"--name", "my-vpc",
		"--project-id", "123e4567-e89b-12d3-a456-426614174000",
		"--region", "us-wdc-1",
		"--subnet-name", "my-subnet",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestNetworkingVPCsAvailabilityUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:vpcs:availability", "update",
		"--vpc-id", "vpc_id",
		"--name", "my-vpc",
		"--subnet-name", "my-subnet",
		"--tag", "production",
		"--tag", "ethereum",
	)
}
