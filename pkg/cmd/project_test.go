// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestProjectsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"projects", "create",
		"--name", "My Project",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestProjectsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"projects", "update",
		"--project-id", "project_id",
		"--name", "My Updated Project",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestProjectsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"projects", "list",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestProjectsDelete(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"projects", "delete",
		"--project-id", "project_id",
	)
}

func TestProjectsGet(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"projects", "get",
		"--project-id", "project_id",
	)
}
