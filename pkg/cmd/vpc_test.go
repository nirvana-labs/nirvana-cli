// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNetworkingVPCsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "networking:vpcs", "create",
			"--api-key", "string",
			"--name", "my-vpc",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--region", "us-wdc-1",
			"--subnet-name", "my-subnet",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-vpc\n" +
			"project_id: 123e4567-e89b-12d3-a456-426614174000\n" +
			"region: us-wdc-1\n" +
			"subnet_name: my-subnet\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "networking:vpcs", "create",
			"--api-key", "string",
		)
	})
}

func TestNetworkingVPCsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "networking:vpcs", "update",
			"--api-key", "string",
			"--vpc-id", "vpc_id",
			"--name", "my-vpc",
			"--subnet-name", "my-subnet",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: my-vpc\n" +
			"subnet_name: my-subnet\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "networking:vpcs", "update",
			"--api-key", "string",
			"--vpc-id", "vpc_id",
		)
	})
}

func TestNetworkingVPCsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "networking:vpcs", "list",
			"--api-key", "string",
			"--max-items", "10",
			"--project-id", "project_id",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestNetworkingVPCsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "networking:vpcs", "delete",
			"--api-key", "string",
			"--vpc-id", "vpc_id",
		)
	})
}

func TestNetworkingVPCsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "networking:vpcs", "get",
			"--api-key", "string",
			"--vpc-id", "vpc_id",
		)
	})
}
