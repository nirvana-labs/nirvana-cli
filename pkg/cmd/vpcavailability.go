// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/stainless-sdks/nirvana-cli/internal/apiquery"
	"github.com/stainless-sdks/nirvana-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var networkingVPCsAvailabilityCreate = cli.Command{
	Name:  "create",
	Usage: "Check if a VPC can be created",
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
	Action:          handleNetworkingVPCsAvailabilityCreate,
	HideHelpCommand: true,
}

var networkingVPCsAvailabilityUpdate = cli.Command{
	Name:  "update",
	Usage: "Check if a VPC can be updated",
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
	Action:          handleNetworkingVPCsAvailabilityUpdate,
	HideHelpCommand: true,
}

func handleNetworkingVPCsAvailabilityCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := networking.VPCAvailabilityNewParams{}

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
	_, err = client.Networking.VPCs.Availability.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:vpcs:availability create", obj, format, transform)
}

func handleNetworkingVPCsAvailabilityUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vpc-id") && len(unusedArgs) > 0 {
		cmd.Set("vpc-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := networking.VPCAvailabilityUpdateParams{}

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
	_, err = client.Networking.VPCs.Availability.Update(
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
	return ShowJSON(os.Stdout, "networking:vpcs:availability update", obj, format, transform)
}
