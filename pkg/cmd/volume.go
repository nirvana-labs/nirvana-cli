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

var computeVolumesCreate = cli.Command{
	Name:  "create",
	Usage: "Create a Volume. Only data volumes can be created.",
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
	Action:          handleComputeVolumesCreate,
	HideHelpCommand: true,
}

var computeVolumesUpdate = cli.Command{
	Name:  "update",
	Usage: "Update a Volume. Boot or data volumes can be updated.",
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
	Action:          handleComputeVolumesUpdate,
	HideHelpCommand: true,
}

var computeVolumesList = cli.Command{
	Name:  "list",
	Usage: "List all volumes",
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
	Action:          handleComputeVolumesList,
	HideHelpCommand: true,
}

var computeVolumesDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete a Volume. Boot or data volumes can be deleted.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "volume-id",
		},
	},
	Action:          handleComputeVolumesDelete,
	HideHelpCommand: true,
}

var computeVolumesGet = cli.Command{
	Name:  "get",
	Usage: "Get a Volume.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "volume-id",
		},
	},
	Action:          handleComputeVolumesGet,
	HideHelpCommand: true,
}

func handleComputeVolumesCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := compute.VolumeNewParams{}

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
	_, err = client.Compute.Volumes.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:volumes create", obj, format, transform)
}

func handleComputeVolumesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("volume-id") && len(unusedArgs) > 0 {
		cmd.Set("volume-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := compute.VolumeUpdateParams{}

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
	_, err = client.Compute.Volumes.Update(
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
	return ShowJSON(os.Stdout, "compute:volumes update", obj, format, transform)
}

func handleComputeVolumesList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := compute.VolumeListParams{}

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
		_, err = client.Compute.Volumes.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "compute:volumes list", obj, format, transform)
	} else {
		iter := client.Compute.Volumes.ListAutoPaging(ctx, params, options...)
		return streamOutput("compute:volumes list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "compute:volumes list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleComputeVolumesDelete(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.Volumes.Delete(ctx, cmd.Value("volume-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:volumes delete", obj, format, transform)
}

func handleComputeVolumesGet(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.Volumes.Get(ctx, cmd.Value("volume-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "compute:volumes get", obj, format, transform)
}
