// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var networkingVPCsCreate = cli.Command{
	Name:  "create",
	Usage: "Create a VPC",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the VPC.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region the resource is in.",
			BodyPath: "region",
		},
		&requestflag.Flag[string]{
			Name:     "subnet-name",
			Usage:    "Name of the subnet to create.",
			BodyPath: "subnet_name",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the VPC.",
			BodyPath: "tags",
		},
	},
	Action:          handleNetworkingVPCsCreate,
	HideHelpCommand: true,
}

var networkingVPCsUpdate = cli.Command{
	Name:  "update",
	Usage: "Update a VPC",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vpc-id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the VPC.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "subnet-name",
			Usage:    "Name of the subnet to create.",
			BodyPath: "subnet_name",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the VPC.",
			BodyPath: "tags",
		},
	},
	Action:          handleNetworkingVPCsUpdate,
	HideHelpCommand: true,
}

var networkingVPCsList = cli.Command{
	Name:  "list",
	Usage: "List all VPCs",
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
	Action:          handleNetworkingVPCsList,
	HideHelpCommand: true,
}

var networkingVPCsDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete a VPC",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vpc-id",
		},
	},
	Action:          handleNetworkingVPCsDelete,
	HideHelpCommand: true,
}

var networkingVPCsGet = cli.Command{
	Name:  "get",
	Usage: "Get details about a VPC",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vpc-id",
		},
	},
	Action:          handleNetworkingVPCsGet,
	HideHelpCommand: true,
}

func handleNetworkingVPCsCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := networking.VPCNewParams{}

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
	_, err = client.Networking.VPCs.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:vpcs create", obj, format, transform)
}

func handleNetworkingVPCsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vpc-id") && len(unusedArgs) > 0 {
		cmd.Set("vpc-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := networking.VPCUpdateParams{}

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
	_, err = client.Networking.VPCs.Update(
		ctx,
		cmd.Value("vpc-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:vpcs update", obj, format, transform)
}

func handleNetworkingVPCsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := networking.VPCListParams{}

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
		_, err = client.Networking.VPCs.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "networking:vpcs list", obj, format, transform)
	} else {
		iter := client.Networking.VPCs.ListAutoPaging(ctx, params, options...)
		return streamOutput("networking:vpcs list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "networking:vpcs list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleNetworkingVPCsDelete(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vpc-id") && len(unusedArgs) > 0 {
		cmd.Set("vpc-id", unusedArgs[0])
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
	_, err = client.Networking.VPCs.Delete(ctx, cmd.Value("vpc-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:vpcs delete", obj, format, transform)
}

func handleNetworkingVPCsGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vpc-id") && len(unusedArgs) > 0 {
		cmd.Set("vpc-id", unusedArgs[0])
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
	_, err = client.Networking.VPCs.Get(ctx, cmd.Value("vpc-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:vpcs get", obj, format, transform)
}
