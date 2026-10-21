// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestOrganizationsMembershipsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:memberships", "update",
			"--organization-id", "organization_id",
			"--membership-id", "membership_id",
			"--role", "owner",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("role: owner")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organizations:memberships", "update",
			"--organization-id", "organization_id",
			"--membership-id", "membership_id",
		)
	})
}

func TestOrganizationsMembershipsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:memberships", "list",
			"--max-items", "10",
			"--organization-id", "organization_id",
			"--cursor", "cursor",
			"--limit", "10",
			"--role", "owner",
			"--sort", "sort",
			"--user-id", "user_id",
		)
	})
}

func TestOrganizationsMembershipsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:memberships", "delete",
			"--organization-id", "organization_id",
			"--membership-id", "membership_id",
		)
	})
}

func TestOrganizationsMembershipsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:memberships", "get",
			"--organization-id", "organization_id",
			"--membership-id", "membership_id",
		)
	})
}
