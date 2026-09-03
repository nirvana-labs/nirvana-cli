// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNKSKubernetesVersionsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:kubernetes-versions", "list",
			"--max-items", "10",
			"--cursor", "cursor",
			"--limit", "10",
			"--name", "name",
			"--sort", "sort",
		)
	})
}
