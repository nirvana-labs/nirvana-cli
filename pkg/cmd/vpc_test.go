// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNetworkingVPCsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:vpcs", "create",
		"--api-key", "string",
		"--name", "my-vpc",
		"--project-id", "123e4567-e89b-12d3-a456-426614174000",
		"--region", "us-wdc-1",
		"--subnet-name", "my-subnet",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestNetworkingVPCsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:vpcs", "update",
		"--api-key", "string",
		"--vpc-id", "vpc_id",
		"--name", "my-vpc",
		"--subnet-name", "my-subnet",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestNetworkingVPCsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:vpcs", "list",
		"--api-key", "string",
		"--project-id", "project_id",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestNetworkingVPCsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:vpcs", "delete",
		"--api-key", "string",
		"--vpc-id", "vpc_id",
	)
}

func TestNetworkingVPCsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:vpcs", "get",
		"--api-key", "string",
		"--vpc-id", "vpc_id",
	)
}
