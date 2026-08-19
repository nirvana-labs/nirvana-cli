// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestAPIKeysCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "create",
			"--expires-at", "'2025-12-31T23:59:59Z'",
			"--name", "My API Key",
			"--permission", "{permission: edit, resource_type: vm}",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--project-id", "123e4567-e89b-12d3-a456-426614174001",
			"--source-ip-rule", "{allowed: [192.168.1.0/24, 10.0.0.0/8], blocked: [192.168.1.100/32]}",
			"--starts-at", "'2025-01-01T00:00:00Z'",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(apiKeysCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "create",
			"--expires-at", "'2025-12-31T23:59:59Z'",
			"--name", "My API Key",
			"--permission.permission", "edit",
			"--permission.resource-type", "vm",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--project-id", "123e4567-e89b-12d3-a456-426614174001",
			"--source-ip-rule.allowed", "[192.168.1.0/24, 10.0.0.0/8]",
			"--source-ip-rule.blocked", "[192.168.1.100/32]",
			"--starts-at", "'2025-01-01T00:00:00Z'",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"expires_at: '2025-12-31T23:59:59Z'\n" +
			"name: My API Key\n" +
			"permissions:\n" +
			"  - permission: edit\n" +
			"    resource_type: vm\n" +
			"project_ids:\n" +
			"  - 123e4567-e89b-12d3-a456-426614174000\n" +
			"  - 123e4567-e89b-12d3-a456-426614174001\n" +
			"source_ip_rule:\n" +
			"  allowed:\n" +
			"    - 192.168.1.0/24\n" +
			"    - 10.0.0.0/8\n" +
			"  blocked:\n" +
			"    - 192.168.1.100/32\n" +
			"starts_at: '2025-01-01T00:00:00Z'\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"api-keys", "create",
		)
	})
}

func TestAPIKeysUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "update",
			"--api-key-id", "api_key_id",
			"--name", "My Updated API Key",
			"--permission", "{permission: edit, resource_type: vm}",
			"--project-id", "string",
			"--source-ip-rule", "{allowed: [192.168.1.0/24, 10.0.0.0/8], blocked: [192.168.1.100/32]}",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(apiKeysUpdate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "update",
			"--api-key-id", "api_key_id",
			"--name", "My Updated API Key",
			"--permission.permission", "edit",
			"--permission.resource-type", "vm",
			"--project-id", "string",
			"--source-ip-rule.allowed", "[192.168.1.0/24, 10.0.0.0/8]",
			"--source-ip-rule.blocked", "[192.168.1.100/32]",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"name: My Updated API Key\n" +
			"permissions:\n" +
			"  - permission: edit\n" +
			"    resource_type: vm\n" +
			"project_ids:\n" +
			"  - string\n" +
			"source_ip_rule:\n" +
			"  allowed:\n" +
			"    - 192.168.1.0/24\n" +
			"    - 10.0.0.0/8\n" +
			"  blocked:\n" +
			"    - 192.168.1.100/32\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"api-keys", "update",
			"--api-key-id", "api_key_id",
		)
	})
}

func TestAPIKeysList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "10",
			"--name", "name",
			"--sort", "sort",
			"--status", "active",
			"--tag", "string",
		)
	})
}

func TestAPIKeysDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "delete",
			"--api-key-id", "api_key_id",
		)
	})
}

func TestAPIKeysGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"api-keys", "get",
			"--api-key-id", "api_key_id",
		)
	})
}
