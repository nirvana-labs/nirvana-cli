// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestOrganizationsBillingRechargePolicyUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:billing:recharge-policy", "update",
			"--organization-id", "organization_id",
			"--policy", "automatic",
			"--policy-args", "{fixed: '-69125', runway_days: '-69125', monthly_cap: '-69125'}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(organizationsBillingRechargePolicyUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:billing:recharge-policy", "update",
			"--organization-id", "organization_id",
			"--policy", "automatic",
			"--policy-args.fixed", "-69125",
			"--policy-args.runway-days", "-69125",
			"--policy-args.monthly-cap", "-69125",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"policy: automatic\n" +
			"policy_args:\n" +
			"  fixed: '-69125'\n" +
			"  runway_days: '-69125'\n" +
			"  monthly_cap: '-69125'\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"organizations:billing:recharge-policy", "update",
			"--organization-id", "organization_id",
		)
	})
}

func TestOrganizationsBillingRechargePolicyGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"organizations:billing:recharge-policy", "get",
			"--organization-id", "organization_id",
		)
	})
}
