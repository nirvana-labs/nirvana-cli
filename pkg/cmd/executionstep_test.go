// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestVektorExecutionsStepsGet(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:executions:steps", "get",
		"--execution-id", "execution_id",
		"--step-id", "step_id",
	)
}

func TestVektorExecutionsStepsSign(t *testing.T) {
	t.Skip("Prism doesn't support callbacks yet")
	mocktest.TestRunMockTestWithFlags(
		t,
		"vektor:executions:steps", "sign",
		"--execution-id", "execution_id",
		"--step-id", "step_id",
		"--signed-payload", "0x123456789abcdef",
	)
}
