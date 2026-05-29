// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestOrganizationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "create",
			"--name", "My Organization",
			"--billing-email", "billing@example.com",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My Organization\n" +
			"billing_email: billing@example.com\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organizations", "create",
		)
	})
}

func TestOrganizationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "update",
			"--organization-id", "organization_id",
			"--billing-email", "billing@example.com",
			"--name", "My Updated Organization",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"billing_email: billing@example.com\n" +
			"name: My Updated Organization\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organizations", "update",
			"--organization-id", "organization_id",
		)
	})
}

func TestOrganizationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestOrganizationsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "get",
			"--organization-id", "organization_id",
		)
	})
}

func TestOrganizationsLeave(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "leave",
			"--organization-id", "organization_id",
		)
	})
}
