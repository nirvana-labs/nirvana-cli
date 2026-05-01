// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/urfave/cli/v3"
)

var computeVolumesAvailabilityCreate = cli.Command{
	Name:    "create",
	Usage:   "Check Volume Create Availability",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Volume.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID the Volume belongs to.",
			Required: true,
			BodyPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region the resource is in.",
			Required: true,
			BodyPath: "region",
		},
		&requestflag.Flag[int64]{
			Name:     "size",
			Usage:    "Size of the Volume in GB.",
			Required: true,
			BodyPath: "size",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Type of the Volume.",
			Required: true,
			BodyPath: "type",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Volume.",
			BodyPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:     "vm-id",
			Usage:    "ID of the VM the Volume is attached to.",
			BodyPath: "vm_id",
		},
	},
	Action:          handleComputeVolumesAvailabilityCreate,
	HideHelpCommand: true,
}

var computeVolumesAvailabilityUpdate = cli.Command{
	Name:    "update",
	Usage:   "Check Volume Update Availability",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "volume-id",
			Required:  true,
			PathParam: "volume_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Volume.",
			BodyPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:     "size",
			Usage:    "Size of the Volume in GB.",
			BodyPath: "size",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Volume.",
			BodyPath: "tags",
		},
	},
	Action:          handleComputeVolumesAvailabilityUpdate,
	HideHelpCommand: true,
}

func handleComputeVolumesAvailabilityCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := compute.VolumeAvailabilityNewParams{}

	return client.Compute.Volumes.Availability.New(ctx, params, options...)
}

func handleComputeVolumesAvailabilityUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("volume-id") && len(unusedArgs) > 0 {
		cmd.Set("volume-id", unusedArgs[0])
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

	params := compute.VolumeAvailabilityUpdateParams{}

	return client.Compute.Volumes.Availability.Update(
		ctx,
		cmd.Value("volume-id").(string),
		params,
		options...,
	)
}
