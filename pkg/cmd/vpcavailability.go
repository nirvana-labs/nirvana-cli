// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/urfave/cli/v3"
)

var networkingVPCsAvailabilityCreate = cli.Command{
	Name:    "create",
	Usage:   "Check if a VPC can be created",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the VPC.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID the VPC belongs to.",
			Required: true,
			BodyPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region the resource is in.",
			Required: true,
			BodyPath: "region",
		},
		&requestflag.Flag[string]{
			Name:     "subnet-name",
			Usage:    "Name of the subnet to create.",
			Required: true,
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
	Name:    "update",
	Usage:   "Check if a VPC can be updated",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "vpc-id",
			Required: true,
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
		false,
	)
	if err != nil {
		return err
	}

	return client.Networking.VPCs.Availability.New(ctx, params, options...)
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
		false,
	)
	if err != nil {
		return err
	}

	return client.Networking.VPCs.Availability.Update(
		ctx,
		cmd.Value("vpc-id").(string),
		params,
		options...,
	)
}
