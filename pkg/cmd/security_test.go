// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestUserSecurityUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"user:security", "update",
		"--source-ip-rule", "{allowed: [192.168.1.0/24, 10.0.0.0/8], blocked: [192.168.1.100/32]}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(userSecurityUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"user:security", "update",
		"--source-ip-rule.allowed", "[192.168.1.0/24, 10.0.0.0/8]",
		"--source-ip-rule.blocked", "[192.168.1.100/32]",
	)
}

func TestUserSecurityGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"user:security", "get",
	)
}
