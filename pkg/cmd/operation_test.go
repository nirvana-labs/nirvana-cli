// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestOperationsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"operations", "list",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestOperationsGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"operations", "get",
		"--operation-id", "operation_id",
	)
}
