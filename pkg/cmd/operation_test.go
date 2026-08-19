// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestOperationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"operations", "list",
			"--max-items", "10",
			"--project-id", "project_id",
			"--created-at-max", "'2019-12-27T18:11:19.117Z'",
			"--created-at-min", "'2019-12-27T18:11:19.117Z'",
			"--cursor", "cursor",
			"--kind", "vm",
			"--limit", "10",
			"--resource-id", "resource_id",
			"--sort", "sort",
			"--status", "pending",
			"--type", "create",
		)
	})
}

func TestOperationsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"operations", "get",
			"--operation-id", "operation_id",
		)
	})
}
