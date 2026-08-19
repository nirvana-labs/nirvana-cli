// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestAuditLogsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"audit-logs", "list",
			"--max-items", "10",
			"--action", "action",
			"--actor-id", "actor_id",
			"--actor-type", "user",
			"--client-ip", "client_ip",
			"--created-at-max", "'2019-12-27T18:11:19.117Z'",
			"--created-at-min", "'2019-12-27T18:11:19.117Z'",
			"--cursor", "cursor",
			"--limit", "10",
			"--method", "method",
			"--path", "path",
			"--sort", "sort",
			"--status-code-max", "0",
			"--status-code-min", "0",
			"--target-id", "target_id",
			"--target-type", "target_type",
		)
	})
}

func TestAuditLogsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"audit-logs", "get",
			"--audit-log-id", "audit_log_id",
		)
	})
}
