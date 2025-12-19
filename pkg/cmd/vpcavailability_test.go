// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNetworkingVPCsAvailabilityCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:vpcs:availability", "create",
		"--name", "my-vpc",
		"--region", "us-wdc-1",
		"--subnet-name", "my-subnet",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestNetworkingVPCsAvailabilityUpdate(t *testing.T) {
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
