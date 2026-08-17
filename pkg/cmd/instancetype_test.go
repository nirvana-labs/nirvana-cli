// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestInstanceTypesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"instance-types", "list",
			"--max-items", "10",
			"--chipset", "chipset",
			"--cursor", "cursor",
			"--family", "family",
			"--limit", "10",
			"--memory-gb-max", "0",
			"--memory-gb-min", "0",
			"--name", "name",
			"--network-bandwidth-gbps-max", "0",
			"--network-bandwidth-gbps-min", "0",
			"--region", "region",
			"--series", "series",
			"--sort", "sort",
			"--vcpu-max", "0",
			"--vcpu-min", "0",
		)
	})
}

func TestInstanceTypesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"instance-types", "get",
			"--region", "us-sva-2",
			"--name", "n1-standard-8",
		)
	})
}
