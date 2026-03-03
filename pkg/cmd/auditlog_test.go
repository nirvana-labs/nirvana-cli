// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestOrganizationsAuditLogsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"organizations:audit-logs", "list",
		"--api-key", "string",
		"--organization-id", "organization_id",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestOrganizationsAuditLogsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"organizations:audit-logs", "get",
		"--api-key", "string",
		"--organization-id", "organization_id",
		"--audit-log-id", "audit_log_id",
	)
}
