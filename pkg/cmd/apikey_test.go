// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestAPIKeysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "create",
		"--api-key", "string",
		"--expires-at", "'2025-12-31T23:59:59Z'",
		"--name", "My API Key",
		"--source-ip-rule", "{allowed: [192.168.1.0/24, 10.0.0.0/8], blocked: [192.168.1.100/32]}",
		"--starts-at", "'2025-01-01T00:00:00Z'",
		"--tag", "production",
		"--tag", "ethereum",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(apiKeysCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "create",
		"--api-key", "string",
		"--expires-at", "'2025-12-31T23:59:59Z'",
		"--name", "My API Key",
		"--source-ip-rule.allowed", "[192.168.1.0/24, 10.0.0.0/8]",
		"--source-ip-rule.blocked", "[192.168.1.100/32]",
		"--starts-at", "'2025-01-01T00:00:00Z'",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestAPIKeysUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "update",
		"--api-key", "string",
		"--api-key-id", "api_key_id",
		"--name", "My Updated API Key",
		"--source-ip-rule", "{allowed: [192.168.1.0/24, 10.0.0.0/8], blocked: [192.168.1.100/32]}",
		"--tag", "production",
		"--tag", "ethereum",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(apiKeysUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "update",
		"--api-key", "string",
		"--api-key-id", "api_key_id",
		"--name", "My Updated API Key",
		"--source-ip-rule.allowed", "[192.168.1.0/24, 10.0.0.0/8]",
		"--source-ip-rule.blocked", "[192.168.1.100/32]",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestAPIKeysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "list",
		"--api-key", "string",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestAPIKeysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "delete",
		"--api-key", "string",
		"--api-key-id", "api_key_id",
	)
}

func TestAPIKeysGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "get",
		"--api-key", "string",
		"--api-key-id", "api_key_id",
	)
}
