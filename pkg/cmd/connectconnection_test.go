// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestNetworkingConnectConnectionsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:connections", "create",
		"--bandwidth-mbps", "50",
		"--cidr", "10.0.0.0/16",
		"--name", "my-connect-connection",
		"--provider-cidr", "172.16.0.0/16",
		"--region", "us-wdc-1",
		"--aws", "{account_id: '523816707215', region: us-east-1}",
		"--project-id", "123e4567-e89b-12d3-a456-426614174000",
		"--tag", "production",
		"--tag", "ethereum",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(networkingConnectConnectionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:connections", "create",
		"--bandwidth-mbps", "50",
		"--cidr", "10.0.0.0/16",
		"--name", "my-connect-connection",
		"--provider-cidr", "172.16.0.0/16",
		"--region", "us-wdc-1",
		"--aws.account-id", "523816707215",
		"--aws.region", "us-east-1",
		"--project-id", "123e4567-e89b-12d3-a456-426614174000",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestNetworkingConnectConnectionsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
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
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:connections", "list",
		"--cursor", "cursor",
		"--limit", "10",
		"--project-id", "project_id",
	)
}

func TestNetworkingConnectConnectionsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:connections", "delete",
		"--connection-id", "connection_id",
	)
}

func TestNetworkingConnectConnectionsGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"networking:connect:connections", "get",
		"--connection-id", "connection_id",
	)
}
