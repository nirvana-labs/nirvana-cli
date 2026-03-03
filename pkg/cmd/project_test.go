// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestProjectsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"projects", "create",
		"--api-key", "string",
		"--name", "My Project",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestProjectsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"projects", "update",
		"--api-key", "string",
		"--project-id", "project_id",
		"--name", "My Updated Project",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestProjectsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"projects", "list",
		"--api-key", "string",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestProjectsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"projects", "delete",
		"--api-key", "string",
		"--project-id", "project_id",
	)
}

func TestProjectsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"projects", "get",
		"--api-key", "string",
		"--project-id", "project_id",
	)
}
