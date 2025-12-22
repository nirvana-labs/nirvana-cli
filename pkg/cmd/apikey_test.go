// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestAPIKeysCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "create",
		"--expires-at", "2025-12-31T23:59:59Z",
		"--name", "My API Key",
		"--source-ip-rule", "{allowed: [192.168.1.0/24, 10.0.0.0/8], blocked: [192.168.1.100/32]}",
		"--starts-at", "2025-01-01T00:00:00Z",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestAPIKeysUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "update",
		"--api-key-id", "api_key_id",
		"--name", "My Updated API Key",
		"--source-ip-rule", "{allowed: [192.168.1.0/24, 10.0.0.0/8], blocked: [192.168.1.100/32]}",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestAPIKeysList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "list",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestAPIKeysDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "delete",
		"--api-key-id", "api_key_id",
	)
}

func TestAPIKeysGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"api-keys", "get",
		"--api-key-id", "api_key_id",
	)
}
