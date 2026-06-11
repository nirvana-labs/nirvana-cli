// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var networkingConnectConnectionsCostCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Return a priced cost quote for the proposed Connect Connection.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "bandwidth-mbps",
			Usage:    "Connect Connection speed in Mbps",
			Required: true,
			BodyPath: "bandwidth_mbps",
		},
		&requestflag.Flag[[]string]{
			Name:     "cidr",
			Usage:    "CIDRs for the Connect Connection. Must be in network-aligned/canonical form.",
			Required: true,
			BodyPath: "cidrs",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Connect Connection",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID the Connect Connection belongs to",
			Required: true,
			BodyPath: "project_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "provider-cidr",
			Usage:    "Provider CIDRs. Must be in network-aligned/canonical form.",
			Required: true,
			BodyPath: "provider_cidrs",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region the resource is in.",
			Required: true,
			BodyPath: "region",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "aws",
			Usage:    "AWS provider configuration",
			BodyPath: "aws",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Connect Connection",
			BodyPath: "tags",
		},
	},
	Action:          handleNetworkingConnectConnectionsCostCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"aws": {
		&requestflag.InnerFlag[string]{
			Name:       "aws.account-id",
			Usage:      "AWS account id",
			InnerField: "account_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aws.region",
			Usage:      "AWS region where the connection will be established",
			InnerField: "region",
		},
	},
})

var networkingConnectConnectionsCostUpdate = cli.Command{
	Name:    "update",
	Usage:   "Return a priced cost quote for the proposed Connect Connection update plus a\ndiff against the current state.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "connection-id",
			Required:  true,
			PathParam: "connection_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Connect Connection.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Connect Connection",
			BodyPath: "tags",
		},
	},
	Action:          handleNetworkingConnectConnectionsCostUpdate,
	HideHelpCommand: true,
}

func handleNetworkingConnectConnectionsCostCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := networking.ConnectConnectionCostNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Networking.Connect.Connections.Cost.New(ctx, params, options...)
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
		Title:          "networking:connect:connections:cost create",
		Transform:      transform,
	})
}

func handleNetworkingConnectConnectionsCostUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
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

	params := networking.ConnectConnectionCostUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Networking.Connect.Connections.Cost.Update(
		ctx,
		cmd.Value("connection-id").(string),
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
		Title:          "networking:connect:connections:cost update",
		Transform:      transform,
	})
}
