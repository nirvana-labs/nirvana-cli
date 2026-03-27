// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestOrganizationsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "create",
			"--name", "My Organization",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("name: My Organization")
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
			"--name", "My Updated Organization",
			"--settings", "{jit_provisioning: false}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(organizationsUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations", "update",
			"--organization-id", "organization_id",
			"--name", "My Updated Organization",
			"--settings.jit-provisioning=false",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My Updated Organization\n" +
			"settings:\n" +
			"  jit_provisioning: false\n")
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
