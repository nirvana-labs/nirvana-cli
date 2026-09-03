// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNKSClustersControllersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:controllers", "list",
			"--max-items", "10",
			"--cluster-id", "cluster_id",
			"--cursor", "cursor",
			"--has-private-ip=true",
			"--instance-type", "instance_type",
			"--limit", "10",
			"--name", "name",
			"--private-ip", "private_ip",
			"--sort", "sort",
			"--status", "ready",
		)
	})
}

func TestNKSClustersControllersGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:controllers", "get",
			"--cluster-id", "cluster_id",
			"--controller-id", "controller_id",
		)
	})
}
