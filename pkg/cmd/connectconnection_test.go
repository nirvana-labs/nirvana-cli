// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestNetworkingConnectConnectionsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:connections", "create",
		"--bandwidth-mbps", "50",
		"--cidr", "10.0.0.0/16",
		"--name", "my-connect-connection",
		"--provider-cidr", "172.16.0.0/16",
		"--region", "us-wdc-1",
		"--aws", "{account_id: '523816707215', region: us-east-1}",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestNetworkingConnectConnectionsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:connections", "update",
		"--connection-id", "connection_id",
		"--name", "my-connect-connection",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestNetworkingConnectConnectionsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:connections", "list",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestNetworkingConnectConnectionsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:connections", "delete",
		"--connection-id", "connection_id",
	)
}

func TestNetworkingConnectConnectionsGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:connections", "get",
		"--connection-id", "connection_id",
	)
}
