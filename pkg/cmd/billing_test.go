// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestOrganizationsBillingCost(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:billing", "cost",
			"--organization-id", "organization_id",
			"--from", "'2019-12-27'",
			"--to", "'2019-12-27'",
		)
	})
}

func TestOrganizationsBillingHistory(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:billing", "history",
			"--organization-id", "organization_id",
			"--created-at-max", "'2019-12-27T18:11:19.117Z'",
			"--created-at-min", "'2019-12-27T18:11:19.117Z'",
			"--currency", "currency",
			"--cursor", "cursor",
			"--limit", "10",
			"--purpose", "purpose",
			"--sort", "sort",
			"--type", "type",
		)
	})
}

func TestOrganizationsBillingRecharge(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:billing", "recharge",
			"--organization-id", "organization_id",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}

func TestOrganizationsBillingStatements(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:billing", "statements",
			"--organization-id", "organization_id",
			"--month", "month",
		)
	})
}

func TestOrganizationsBillingSummary(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:billing", "summary",
			"--organization-id", "organization_id",
		)
	})
}

func TestOrganizationsBillingTopUp(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:billing", "top-up",
			"--organization-id", "organization_id",
			"--amount", "50.00",
			"--idempotency-key", "Idempotency-Key",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("amount: '50.00'")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organizations:billing", "top-up",
			"--organization-id", "organization_id",
			"--idempotency-key", "Idempotency-Key",
		)
	})
}
