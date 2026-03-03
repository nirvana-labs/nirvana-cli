// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNetworkingFirewallRulesCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:firewall-rules", "create",
		"--api-key", "string",
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
}

func TestNetworkingFirewallRulesUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:firewall-rules", "update",
		"--api-key", "string",
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
}

func TestNetworkingFirewallRulesList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:firewall-rules", "list",
		"--api-key", "string",
		"--vpc-id", "vpc_id",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestNetworkingFirewallRulesDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:firewall-rules", "delete",
		"--api-key", "string",
		"--vpc-id", "vpc_id",
		"--firewall-rule-id", "firewall_rule_id",
	)
}

func TestNetworkingFirewallRulesGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:firewall-rules", "get",
		"--api-key", "string",
		"--vpc-id", "vpc_id",
		"--firewall-rule-id", "firewall_rule_id",
	)
}
