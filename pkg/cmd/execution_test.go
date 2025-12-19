// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorExecutionsList(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:executions", "list",
	)
}

func TestVektorExecutionsGet(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:executions", "get",
		"--execution-id", "execution_id",
	)
}
