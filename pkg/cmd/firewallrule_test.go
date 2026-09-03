// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNetworkingFirewallRulesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:firewall-rules", "create",
			"--vpc-id", "vpc_id",
			"--destination-address", "10.0.0.0/25",
			"--destination-port", "22",
			"--destination-port", "80",
			"--destination-port", "443",
			"--name", "my-firewall-rule",
			"--protocol", "tcp",
			"--source-address", "0.0.0.0/0",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destination_address: 10.0.0.0/25\n" +
			"destination_ports:\n" +
			"  - '22'\n" +
			"  - '80'\n" +
			"  - '443'\n" +
			"name: my-firewall-rule\n" +
			"protocol: tcp\n" +
			"source_address: 0.0.0.0/0\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"networking:firewall-rules", "create",
			"--vpc-id", "vpc_id",
		)
	})
}

func TestNetworkingFirewallRulesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:firewall-rules", "update",
			"--vpc-id", "vpc_id",
			"--firewall-rule-id", "firewall_rule_id",
			"--destination-address", "10.0.0.0/25",
			"--destination-port", "22",
			"--destination-port", "80",
			"--destination-port", "443",
			"--name", "my-firewall-rule",
			"--protocol", "tcp",
			"--source-address", "0.0.0.0/0",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"destination_address: 10.0.0.0/25\n" +
			"destination_ports:\n" +
			"  - '22'\n" +
			"  - '80'\n" +
			"  - '443'\n" +
			"name: my-firewall-rule\n" +
			"protocol: tcp\n" +
			"source_address: 0.0.0.0/0\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"networking:firewall-rules", "update",
			"--vpc-id", "vpc_id",
			"--firewall-rule-id", "firewall_rule_id",
		)
	})
}

func TestNetworkingFirewallRulesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:firewall-rules", "list",
			"--max-items", "10",
			"--vpc-id", "vpc_id",
			"--cursor", "cursor",
			"--limit", "10",
			"--name", "name",
			"--protocol", "tcp",
			"--sort", "sort",
			"--status", "ready",
			"--tag", "string",
		)
	})
}

func TestNetworkingFirewallRulesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:firewall-rules", "delete",
			"--vpc-id", "vpc_id",
			"--firewall-rule-id", "firewall_rule_id",
		)
	})
}

func TestNetworkingFirewallRulesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"networking:firewall-rules", "get",
			"--vpc-id", "vpc_id",
			"--firewall-rule-id", "firewall_rule_id",
		)
	})
}
