// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestUserGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"user", "get",
	)
}
