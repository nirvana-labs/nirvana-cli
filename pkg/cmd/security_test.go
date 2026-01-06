// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestUserSecurityUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"user:security", "update",
		"--source-ip-rule", "{allowed: [192.168.1.0/24, 10.0.0.0/8], blocked: [192.168.1.100/32]}",
	)
}

func TestUserSecurityGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"user:security", "get",
	)
}
