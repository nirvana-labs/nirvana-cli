// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNKSClustersPersistentVolumeClaimsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:persistent-volume-claims", "list",
			"--max-items", "10",
			"--cluster-id", "cluster_id",
			"--cursor", "cursor",
			"--limit", "10",
			"--name", "name",
			"--size-gb-max", "0",
			"--size-gb-min", "0",
			"--sort", "sort",
			"--status", "ready",
			"--type", "abs",
		)
	})
}

func TestNKSClustersPersistentVolumeClaimsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:persistent-volume-claims", "get",
			"--cluster-id", "cluster_id",
			"--persistent-volume-claim-id", "persistent_volume_claim_id",
		)
	})
}
