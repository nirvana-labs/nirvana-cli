// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestOperationsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"operations", "list",
		"--api-key", "string",
		"--project-id", "project_id",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestOperationsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"operations", "get",
		"--api-key", "string",
		"--operation-id", "operation_id",
	)
}
