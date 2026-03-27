// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNKSClustersLoadBalancersList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:load-balancers", "list",
			"--max-items", "10",
			"--cluster-id", "cluster_id",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestNKSClustersLoadBalancersGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"nks:clusters:load-balancers", "get",
			"--cluster-id", "cluster_id",
			"--load-balancer-id", "load_balancer_id",
		)
	})
}
