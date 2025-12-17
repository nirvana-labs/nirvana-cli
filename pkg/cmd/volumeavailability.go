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

var computeVolumesAvailabilityCreate = cli.Command{
	Name:  "create",
	Usage: "Check Volume Create Availability",
	Flags: []cli.Flag{
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
		&requestflag.Flag[string]{
			Name:     "type",
			Usage:    "Type of the Volume.",
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "vm-id",
			Usage:    "ID of the VM the Volume is attached to.",
			BodyPath: "vm_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Volume.",
			BodyPath: "tags",
		},
	},
	Action:          handleComputeVolumesAvailabilityCreate,
	HideHelpCommand: true,
}

var computeVolumesAvailabilityUpdate = cli.Command{
	Name:  "update",
	Usage: "Check Volume Update Availability",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "volume-id",
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

	params := compute.VolumeAvailabilityNewParams{}

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
	_, err = client.Compute.Volumes.Availability.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:volumes:availability create", obj, format, transform)
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

	params := compute.VolumeAvailabilityUpdateParams{}

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
	_, err = client.Compute.Volumes.Availability.Update(
		ctx,
		cmd.Value("volume-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:volumes:availability update", obj, format, transform)
}
