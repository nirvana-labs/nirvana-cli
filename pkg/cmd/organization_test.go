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
			t, "organizations", "create",
			"--api-key", "string",
			"--name", "My Organization",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: My Organization")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "organizations", "create",
			"--api-key", "string",
		)
	})
}

func TestOrganizationsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "organizations", "update",
			"--api-key", "string",
			"--organization-id", "organization_id",
			"--name", "My Updated Organization",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: My Updated Organization")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData, "organizations", "update",
			"--api-key", "string",
			"--organization-id", "organization_id",
		)
	})
}

func TestOrganizationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "organizations", "list",
			"--api-key", "string",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestOrganizationsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t, "organizations", "get",
			"--api-key", "string",
			"--organization-id", "organization_id",
		)
	})
}
