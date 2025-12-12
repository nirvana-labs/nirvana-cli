// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/stainless-sdks/nirvana-cli/internal/apiquery"
	"github.com/stainless-sdks/nirvana-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var computeVMsAvailabilityCreate = cli.Command{
	Name:  "create",
	Usage: "Check VM Create Availability",
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "boot-volume",
			Usage:    "Boot volume for the VM.",
			BodyPath: "boot_volume",
		},
		&requestflag.Flag[any]{
			Name:     "cpu-config",
			Usage:    "CPU configuration for the VM.",
			BodyPath: "cpu_config",
		},
		&requestflag.Flag[any]{
			Name:     "memory-config",
			Usage:    "Memory configuration for the VM.",
			BodyPath: "memory_config",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the VM.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "os-image-name",
			Usage:    "Name of the OS Image to use for the VM.",
			BodyPath: "os_image_name",
		},
		&requestflag.Flag[bool]{
			Name:     "public-ip-enabled",
			Usage:    "Whether to enable public IP for the VM.",
			BodyPath: "public_ip_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region the resource is in.",
			BodyPath: "region",
		},
		&requestflag.Flag[any]{
			Name:     "ssh-key",
			Usage:    "Public SSH key configuration for the VM.",
			BodyPath: "ssh_key",
		},
		&requestflag.Flag[string]{
			Name:     "subnet-id",
			Usage:    "ID of the subnet to use for the VM.",
			BodyPath: "subnet_id",
		},
		&requestflag.Flag[[]any]{
			Name:     "data-volume",
			Usage:    "Data volumes for the VM.",
			BodyPath: "data_volumes",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the VM.",
			BodyPath: "tags",
		},
	},
	Action:          handleComputeVMsAvailabilityCreate,
	HideHelpCommand: true,
}

var computeVMsAvailabilityUpdate = cli.Command{
	Name:  "update",
	Usage: "Check VM Update Availability",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vm-id",
		},
		&requestflag.Flag[any]{
			Name:     "cpu-config",
			Usage:    "CPU configuration for the VM.",
			BodyPath: "cpu_config",
		},
		&requestflag.Flag[any]{
			Name:     "memory-config",
			Usage:    "Memory configuration for the VM.",
			BodyPath: "memory_config",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the VM.",
			BodyPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:     "public-ip-enabled",
			Usage:    "Whether to enable public IP for the VM.",
			BodyPath: "public_ip_enabled",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the VM.",
			BodyPath: "tags",
		},
	},
	Action:          handleComputeVMsAvailabilityUpdate,
	HideHelpCommand: true,
}

func handleComputeVMsAvailabilityCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := compute.VMAvailabilityNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.Availability.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:vms:availability create", obj, format, transform)
}

func handleComputeVMsAvailabilityUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vm-id") && len(unusedArgs) > 0 {
		cmd.Set("vm-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := compute.VMAvailabilityUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.Availability.Update(
		ctx,
		cmd.Value("vm-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:vms:availability update", obj, format, transform)
}
