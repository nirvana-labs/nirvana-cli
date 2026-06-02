// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestOrganizationsAddressCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:address", "create",
			"--organization-id", "organization_id",
			"--city", "San Francisco",
			"--country", "US",
			"--line1", "123 Main St",
			"--postal-code", "94105",
			"--line2", "Suite 400",
			"--state", "CA",
			"--tax-id", "EU372000000",
			"--tax-id-type", "eu_vat",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"city: San Francisco\n" +
			"country: US\n" +
			"line1: 123 Main St\n" +
			"postal_code: '94105'\n" +
			"line2: Suite 400\n" +
			"state: CA\n" +
			"tax_id: EU372000000\n" +
			"tax_id_type: eu_vat\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organizations:address", "create",
			"--organization-id", "organization_id",
		)
	})
}

func TestOrganizationsAddressUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:address", "update",
			"--organization-id", "organization_id",
			"--city", "San Francisco",
			"--country", "US",
			"--line1", "123 Main St",
			"--line2", "Suite 400",
			"--postal-code", "94105",
			"--state", "CA",
			"--tax-id", "EU372000000",
			"--tax-id-type", "eu_vat",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"city: San Francisco\n" +
			"country: US\n" +
			"line1: 123 Main St\n" +
			"line2: Suite 400\n" +
			"postal_code: '94105'\n" +
			"state: CA\n" +
			"tax_id: EU372000000\n" +
			"tax_id_type: eu_vat\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organizations:address", "update",
			"--organization-id", "organization_id",
		)
	})
}

func TestOrganizationsAddressGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:address", "get",
			"--organization-id", "organization_id",
		)
	})
}
