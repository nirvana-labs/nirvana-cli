// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNetworkingFirewallRulesCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
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
}

func TestNetworkingFirewallRulesUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
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
}

func TestNetworkingFirewallRulesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:firewall-rules", "list",
		"--vpc-id", "vpc_id",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestNetworkingFirewallRulesDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:firewall-rules", "delete",
		"--vpc-id", "vpc_id",
		"--firewall-rule-id", "firewall_rule_id",
	)
}

func TestNetworkingFirewallRulesGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:firewall-rules", "get",
		"--vpc-id", "vpc_id",
		"--firewall-rule-id", "firewall_rule_id",
	)
}
