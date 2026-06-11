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

var computeVMsCostCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Return a priced cost quote for the proposed VM.",
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
	Action:          handleComputeVMsCostCreate,
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

var computeVMsCostUpdate = cli.Command{
	Name:    "update",
	Usage:   "Return a priced cost quote for the proposed VM update plus a diff against the\ncurrent state.",
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
	Action:          handleComputeVMsCostUpdate,
	HideHelpCommand: true,
}

func handleComputeVMsCostCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := compute.VMCostNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.Cost.New(ctx, params, options...)
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
		Title:          "compute:vms:cost create",
		Transform:      transform,
	})
}

func handleComputeVMsCostUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := compute.VMCostUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.Cost.Update(
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
		Title:          "compute:vms:cost update",
		Transform:      transform,
	})
}
