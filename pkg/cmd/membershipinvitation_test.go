// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestOrganizationsMembershipsInvitationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:memberships:invitations", "create",
			"--organization-id", "organization_id",
			"--email", "jane@example.com",
			"--role", "owner",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"email: jane@example.com\n" +
			"role: owner\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organizations:memberships:invitations", "create",
			"--organization-id", "organization_id",
		)
	})
}

func TestOrganizationsMembershipsInvitationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:memberships:invitations", "list",
			"--organization-id", "organization_id",
		)
	})
}

func TestOrganizationsMembershipsInvitationsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:memberships:invitations", "delete",
			"--organization-id", "organization_id",
			"--invitation-id", "invitation_id",
		)
	})
}

func TestOrganizationsMembershipsInvitationsResend(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:memberships:invitations", "resend",
			"--organization-id", "organization_id",
			"--invitation-id", "invitation_id",
		)
	})
}
