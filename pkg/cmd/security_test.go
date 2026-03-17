// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestUserSecurityUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user:security", "update",
			"--source-ip-rule", "{allowed: [192.168.1.0/24, 10.0.0.0/8], blocked: [192.168.1.100/32]}",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(userSecurityUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user:security", "update",
			"--source-ip-rule.allowed", "[192.168.1.0/24, 10.0.0.0/8]",
			"--source-ip-rule.blocked", "[192.168.1.100/32]",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"source_ip_rule:\n" +
			"  allowed:\n" +
			"    - 192.168.1.0/24\n" +
			"    - 10.0.0.0/8\n" +
			"  blocked:\n" +
			"    - 192.168.1.100/32\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"user:security", "update",
		)
	})
}

func TestUserSecurityGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"user:security", "get",
		)
	})
}
