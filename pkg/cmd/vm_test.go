// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestComputeVMsCreate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms", "create",
		"--boot-volume", "{size: 100, type: nvme, tags: [production, ethereum]}",
		"--cpu-config", "{vcpu: 2}",
		"--memory-config", "{size: 2}",
		"--name", "my-vm",
		"--os-image-name", "ubuntu-noble-2025-10-01",
		"--project-id", "123e4567-e89b-12d3-a456-426614174000",
		"--public-ip-enabled=true",
		"--region", "us-wdc-1",
		"--ssh-key", "{public_key: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2}",
		"--subnet-id", "123e4567-e89b-12d3-a456-426614174000",
		"--data-volume", "{name: my-data-volume, size: 100, type: nvme, tags: [production, ethereum]}",
		"--tag", "production",
		"--tag", "ethereum",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(computeVMsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms", "create",
		"--boot-volume.size", "100",
		"--boot-volume.type", "nvme",
		"--boot-volume.tags", "[production, ethereum]",
		"--cpu-config.vcpu", "2",
		"--memory-config.size", "2",
		"--name", "my-vm",
		"--os-image-name", "ubuntu-noble-2025-10-01",
		"--project-id", "123e4567-e89b-12d3-a456-426614174000",
		"--public-ip-enabled=true",
		"--region", "us-wdc-1",
		"--ssh-key.public-key", "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2",
		"--subnet-id", "123e4567-e89b-12d3-a456-426614174000",
		"--data-volume.name", "my-data-volume",
		"--data-volume.size", "100",
		"--data-volume.type", "nvme",
		"--data-volume.tags", "[production, ethereum]",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestComputeVMsUpdate(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms", "update",
		"--vm-id", "vm_id",
		"--cpu-config", "{vcpu: 2}",
		"--memory-config", "{size: 2}",
		"--name", "my-vm",
		"--public-ip-enabled=true",
		"--tag", "production",
		"--tag", "ethereum",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(computeVMsUpdate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms", "update",
		"--vm-id", "vm_id",
		"--cpu-config.vcpu", "2",
		"--memory-config.size", "2",
		"--name", "my-vm",
		"--public-ip-enabled=true",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestComputeVMsList(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms", "list",
		"--project-id", "project_id",
		"--cursor", "cursor",
		"--limit", "10",
	)
}

func TestComputeVMsDelete(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms", "delete",
		"--vm-id", "vm_id",
	)
}

func TestComputeVMsGet(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms", "get",
		"--vm-id", "vm_id",
	)
}

func TestComputeVMsRestart(t *testing.T) {
	t.Skip("Prism tests are disabled")
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms", "restart",
		"--vm-id", "vm_id",
	)
}
