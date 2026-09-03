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
			t,
			"--api-key", "string",
			"networking:vpcs", "create",
			"--name", "my-vpc",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--region", "us-sva-2",
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
			"region: us-sva-2\n" +
			"subnet_name: my-subnet\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"networking:vpcs", "create",
		)
	})
}

func TestNetworkingVPCsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:vpcs", "update",
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
			t, pipeData,
			"--api-key", "string",
			"networking:vpcs", "update",
			"--vpc-id", "vpc_id",
		)
	})
}

func TestNetworkingVPCsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:vpcs", "list",
			"--max-items", "10",
			"--project-id", "project_id",
			"--cursor", "cursor",
			"--limit", "10",
			"--name", "name",
			"--region", "region",
			"--sort", "sort",
			"--status", "ready",
			"--tag", "string",
		)
	})
}

func TestNetworkingVPCsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:vpcs", "delete",
			"--vpc-id", "vpc_id",
		)
	})
}

func TestNetworkingVPCsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:vpcs", "get",
			"--vpc-id", "vpc_id",
		)
	})
}
