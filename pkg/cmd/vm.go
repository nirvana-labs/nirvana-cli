// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var computeVMsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a VM",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "boot-volume",
			Usage:    "Boot volume for the VM.",
			Required: true,
			BodyPath: "boot_volume",
		},
		&requestflag.Flag[string]{
			Name:     "instance-type",
			Usage:    "Instance type name.",
			Required: true,
			BodyPath: "instance_type",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the VM.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "os-image-name",
			Usage:    "Name of the OS Image to use for the VM.",
			Required: true,
			BodyPath: "os_image_name",
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID to create the VM in.",
			Required: true,
			BodyPath: "project_id",
		},
		&requestflag.Flag[bool]{
			Name:     "public-ip-enabled",
			Usage:    "Whether to enable public IP for the VM.",
			Required: true,
			BodyPath: "public_ip_enabled",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region the resource is in.",
			Required: true,
			BodyPath: "region",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "ssh-key",
			Usage:    "Public SSH key configuration for the VM.",
			Required: true,
			BodyPath: "ssh_key",
		},
		&requestflag.Flag[string]{
			Name:     "subnet-id",
			Usage:    "ID of the subnet to use for the VM.",
			Required: true,
			BodyPath: "subnet_id",
		},
		&requestflag.Flag[[]map[string]any]{
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
}, map[string][]requestflag.HasOuterFlag{
	"boot-volume": {
		&requestflag.InnerFlag[int64]{
			Name:       "boot-volume.size",
			Usage:      "Size of the Volume in GB.",
			InnerField: "size",
		},
		&requestflag.InnerFlag[string]{
			Name:       "boot-volume.type",
			Usage:      "Type of the Volume.",
			InnerField: "type",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "boot-volume.tags",
			Usage:      "Tags to attach to the Volume.",
			InnerField: "tags",
		},
	},
	"ssh-key": {
		&requestflag.InnerFlag[string]{
			Name:       "ssh-key.public-key",
			Usage:      "Public key to and use to access the VM.",
			InnerField: "public_key",
		},
	},
	"data-volume": {
		&requestflag.InnerFlag[string]{
			Name:       "data-volume.name",
			Usage:      "Name of the Volume.",
			InnerField: "name",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "data-volume.size",
			Usage:      "Size of the Volume in GB.",
			InnerField: "size",
		},
		&requestflag.InnerFlag[string]{
			Name:       "data-volume.type",
			Usage:      "Type of the Volume.",
			InnerField: "type",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "data-volume.tags",
			Usage:      "Tags to attach to the Volume.",
			InnerField: "tags",
		},
	},
})

var computeVMsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a VM",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "vm-id",
			Required:  true,
			PathParam: "vm_id",
		},
		&requestflag.Flag[string]{
			Name:     "instance-type",
			Usage:    "Instance type name.",
			BodyPath: "instance_type",
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
	Name:    "list",
	Usage:   "List all VMs",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project ID of resources to request",
			Required:  true,
			QueryPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor returned by a previous request. Only valid for the same filters and sort order.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Filter by a case-insensitive substring of the VM name",
			QueryPath: "name",
		},
		&requestflag.Flag[bool]{
			Name:      "public-ip-enabled",
			Usage:     "Filter by whether a public IP is enabled",
			QueryPath: "public_ip_enabled",
		},
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Filter by region",
			QueryPath: "region",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, updated_at, name, status, vcpu, memory",
			Default:   "created_at:desc",
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by VM status",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "subnet-id",
			Usage:     "Filter by the subnet the VM is attached to",
			QueryPath: "subnet_id",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Filter by tags. Repeat the parameter to require several tags; a VM must carry all of them.",
			QueryPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:      "vpc-id",
			Usage:     "Filter by the VPC the VM is attached to",
			QueryPath: "vpc_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleComputeVMsList,
	HideHelpCommand: true,
}

var computeVMsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a VM",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "vm-id",
			Required:  true,
			PathParam: "vm_id",
		},
	},
	Action:          handleComputeVMsDelete,
	HideHelpCommand: true,
}

var computeVMsGet = cli.Command{
	Name:    "get",
	Usage:   "Get details about a VM",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "vm-id",
			Required:  true,
			PathParam: "vm_id",
		},
	},
	Action:          handleComputeVMsGet,
	HideHelpCommand: true,
}

var computeVMsRestart = cli.Command{
	Name:    "restart",
	Usage:   "Restart a VM",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "vm-id",
			Required:  true,
			PathParam: "vm_id",
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

	params := compute.VMNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "compute:vms create",
		Transform:      transform,
	})
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

	params := compute.VMUpdateParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "compute:vms update",
		Transform:      transform,
	})
}

func handleComputeVMsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := compute.VMListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Compute.VMs.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "compute:vms list",
			Transform:      transform,
		})
	} else {
		iter := client.Compute.VMs.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "compute:vms list",
			Transform:      transform,
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "compute:vms delete",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "compute:vms get",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "compute:vms restart",
		Transform:      transform,
	})
}
