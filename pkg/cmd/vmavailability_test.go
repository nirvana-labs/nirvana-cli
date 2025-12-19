// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
)

func TestComputeVMsAvailabilityCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms:availability", "create",
		"--boot-volume", "{size: 100, type: nvme, tags: [production, ethereum]}",
		"--cpu-config", "{vcpu: 2}",
		"--memory-config", "{size: 2}",
		"--name", "my-vm",
		"--os-image-name", "ubuntu-noble-2025-10-01",
		"--public-ip-enabled",
		"--region", "us-wdc-1",
		"--ssh-key", "{public_key: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2}",
		"--subnet-id", "123e4567-e89b-12d3-a456-426614174000",
		"--data-volume", "{name: my-data-volume, size: 100, type: nvme, tags: [production, ethereum]}\n",
		"--tag", "production",
		"--tag", "ethereum",
	)
}

func TestComputeVMsAvailabilityUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"compute:vms:availability", "update",
		"--vm-id", "vm_id",
		"--cpu-config", "{vcpu: 2}",
		"--memory-config", "{size: 2}",
		"--name", "my-vm",
		"--public-ip-enabled",
		"--tag", "production",
		"--tag", "ethereum",
	)
}
