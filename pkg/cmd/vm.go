// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var computeVMsCreate = cli.Command{
	Name:  "create",
	Usage: "Create a VM",
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
	Action:          handleComputeVMsCreate,
	HideHelpCommand: true,
}

var computeVMsUpdate = cli.Command{
	Name:  "update",
	Usage: "Update a VM",
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
	Action:          handleComputeVMsUpdate,
	HideHelpCommand: true,
}

var computeVMsList = cli.Command{
	Name:  "list",
	Usage: "List all VMs",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor returned by a previous request",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
	},
	Action:          handleComputeVMsList,
	HideHelpCommand: true,
}

var computeVMsDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete a VM",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vm-id",
		},
	},
	Action:          handleComputeVMsDelete,
	HideHelpCommand: true,
}

var computeVMsGet = cli.Command{
	Name:  "get",
	Usage: "Get details about a VM",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vm-id",
		},
	},
	Action:          handleComputeVMsGet,
	HideHelpCommand: true,
}

var computeVMsRestart = cli.Command{
	Name:  "restart",
	Usage: "Restart a VM",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vm-id",
		},
	},
	Action:          handleComputeVMsRestart,
	HideHelpCommand: true,
}

func handleComputeVMsCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := compute.VMNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:vms create", obj, format, transform)
}

func handleComputeVMsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vm-id") && len(unusedArgs) > 0 {
		cmd.Set("vm-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := compute.VMUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.Update(
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
	return ShowJSON(os.Stdout, "compute:vms update", obj, format, transform)
}

func handleComputeVMsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := compute.VMListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Compute.VMs.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "compute:vms list", obj, format, transform)
	} else {
		iter := client.Compute.VMs.ListAutoPaging(ctx, params, options...)
		return streamOutput("compute:vms list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "compute:vms list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleComputeVMsDelete(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vm-id") && len(unusedArgs) > 0 {
		cmd.Set("vm-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.Delete(ctx, cmd.Value("vm-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:vms delete", obj, format, transform)
}

func handleComputeVMsGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vm-id") && len(unusedArgs) > 0 {
		cmd.Set("vm-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.Get(ctx, cmd.Value("vm-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:vms get", obj, format, transform)
}

func handleComputeVMsRestart(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vm-id") && len(unusedArgs) > 0 {
		cmd.Set("vm-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.Restart(ctx, cmd.Value("vm-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:vms restart", obj, format, transform)
}
