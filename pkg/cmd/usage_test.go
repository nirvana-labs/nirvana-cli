// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestUsageList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"usage", "list",
			"--max-items", "10",
			"--active-at-max", "'2019-12-27T18:11:19.117Z'",
			"--active-at-min", "'2019-12-27T18:11:19.117Z'",
			"--cursor", "cursor",
			"--dimension", "dimension",
			"--limit", "10",
			"--region", "region",
			"--resource-id", "resource_id",
			"--resource-type", "vm",
			"--sort", "sort",
		)
	})
}

func TestUsageGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"usage", "get",
			"--resource-id", "123e4567-e89b-12d3-a456-426614174000",
		)
	})
}
