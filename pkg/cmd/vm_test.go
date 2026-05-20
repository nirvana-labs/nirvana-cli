// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/nirvana-labs/nirvana-cli/internal/mocktest"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
)

func TestComputeVMsCreate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:vms", "create",
			"--boot-volume", "{size: 100, type: abs, tags: [production, ethereum]}",
			"--instance-type", "n1-standard-8",
			"--name", "my-vm",
			"--os-image-name", "ubuntu-noble-2025-10-01",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--public-ip-enabled=true",
			"--region", "us-sva-2",
			"--ssh-key", "{public_key: ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2}",
			"--subnet-id", "123e4567-e89b-12d3-a456-426614174000",
			"--data-volume", "{name: my-data-volume, size: 100, type: abs, tags: [production, ethereum]}",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("inner flags", func(t *testing.T) {
		// Check that inner flags have been set up correctly
		requestflag.CheckInnerFlags(computeVMsCreate)

		// Alternative argument passing style using inner flags
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:vms", "create",
			"--boot-volume.size", "100",
			"--boot-volume.type", "abs",
			"--boot-volume.tags", "[production, ethereum]",
			"--instance-type", "n1-standard-8",
			"--name", "my-vm",
			"--os-image-name", "ubuntu-noble-2025-10-01",
			"--project-id", "123e4567-e89b-12d3-a456-426614174000",
			"--public-ip-enabled=true",
			"--region", "us-sva-2",
			"--ssh-key.public-key", "ssh-ed25519 AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2",
			"--subnet-id", "123e4567-e89b-12d3-a456-426614174000",
			"--data-volume.name", "my-data-volume",
			"--data-volume.size", "100",
			"--data-volume.type", "abs",
			"--data-volume.tags", "[production, ethereum]",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"boot_volume:\n" +
			"  size: 100\n" +
			"  type: abs\n" +
			"  tags:\n" +
			"    - production\n" +
			"    - ethereum\n" +
			"instance_type: n1-standard-8\n" +
			"name: my-vm\n" +
			"os_image_name: ubuntu-noble-2025-10-01\n" +
			"project_id: 123e4567-e89b-12d3-a456-426614174000\n" +
			"public_ip_enabled: true\n" +
			"region: us-sva-2\n" +
			"ssh_key:\n" +
			"  public_key: >-\n" +
			"    ssh-ed25519\n" +
			"    AAAAC3NzaC1lZDI1NTE5AAAAIDBIASkmwNiLcdlW6927Zjt1Hf7Kw/PpEZ4Zm+wU9wn2\n" +
			"subnet_id: 123e4567-e89b-12d3-a456-426614174000\n" +
			"data_volumes:\n" +
			"  - name: my-data-volume\n" +
			"    size: 100\n" +
			"    type: abs\n" +
			"    tags:\n" +
			"      - production\n" +
			"      - ethereum\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"compute:vms", "create",
		)
	})
}

func TestComputeVMsUpdate(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:vms", "update",
			"--vm-id", "vm_id",
			"--instance-type", "n1-standard-8",
			"--name", "my-vm",
			"--public-ip-enabled=true",
			"--tag", "production",
			"--tag", "ethereum",
		)
	})

	t.Run("piping data", func(t *testing.T) {
		// Test piping YAML data over stdin
		pipeData := []byte("" +
			"instance_type: n1-standard-8\n" +
			"name: my-vm\n" +
			"public_ip_enabled: true\n" +
			"tags:\n" +
			"  - production\n" +
			"  - ethereum\n")
		mocktest.TestRunMockTestWithPipeAndFlags(
			t, pipeData,
			"--api-key", "string",
			"compute:vms", "update",
			"--vm-id", "vm_id",
		)
	})
}

func TestComputeVMsList(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:vms", "list",
			"--max-items", "10",
			"--project-id", "project_id",
			"--cursor", "cursor",
			"--limit", "10",
		)
	})
}

func TestComputeVMsDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:vms", "delete",
			"--vm-id", "vm_id",
		)
	})
}

func TestComputeVMsGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:vms", "get",
			"--vm-id", "vm_id",
		)
	})
}

func TestComputeVMsRestart(t *testing.T) {
	t.Skip("Mock server tests are disabled")
	t.Run("regular flags", func(t *testing.T) {
		mocktest.TestRunMockTestWithFlags(
			t,
			"--api-key", "string",
			"compute:vms", "restart",
			"--vm-id", "vm_id",
		)
	})
}
