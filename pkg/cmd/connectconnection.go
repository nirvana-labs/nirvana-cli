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

var networkingConnectConnectionsCreate = cli.Command{
	Name:  "create",
	Usage: "Create a Connect Connection",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "bandwidth-mbps",
			Usage:    "Connect Connection speed in Mbps",
			BodyPath: "bandwidth_mbps",
		},
		&requestflag.Flag[[]string]{
			Name:     "cidr",
			Usage:    "CIDRs for the Connect Connection. Must be in network-aligned/canonical form.",
			BodyPath: "cidrs",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Connect Connection",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "provider-cidr",
			Usage:    "Provider CIDRs. Must be in network-aligned/canonical form.",
			BodyPath: "provider_cidrs",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region the resource is in.",
			BodyPath: "region",
		},
		&requestflag.Flag[any]{
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
	Action:          handleNetworkingConnectConnectionsCreate,
	HideHelpCommand: true,
}

var networkingConnectConnectionsUpdate = cli.Command{
	Name:  "update",
	Usage: "Update Connect Connection details",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "connection-id",
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
	Action:          handleNetworkingConnectConnectionsUpdate,
	HideHelpCommand: true,
}

var networkingConnectConnectionsList = cli.Command{
	Name:  "list",
	Usage: "List all Connect Connections",
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
	Action:          handleNetworkingConnectConnectionsList,
	HideHelpCommand: true,
}

var networkingConnectConnectionsDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete Connect Connection",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "connection-id",
		},
	},
	Action:          handleNetworkingConnectConnectionsDelete,
	HideHelpCommand: true,
}

var networkingConnectConnectionsGet = cli.Command{
	Name:  "get",
	Usage: "Get Connect Connection details",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "connection-id",
		},
	},
	Action:          handleNetworkingConnectConnectionsGet,
	HideHelpCommand: true,
}

func handleNetworkingConnectConnectionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := networking.ConnectConnectionNewParams{}

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
	_, err = client.Networking.Connect.Connections.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:connect:connections create", obj, format, transform)
}

func handleNetworkingConnectConnectionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("connection-id") && len(unusedArgs) > 0 {
		cmd.Set("connection-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := networking.ConnectConnectionUpdateParams{}

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
	_, err = client.Networking.Connect.Connections.Update(
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:connect:connections update", obj, format, transform)
}

func handleNetworkingConnectConnectionsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := networking.ConnectConnectionListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Networking.Connect.Connections.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "networking:connect:connections list", obj, format, transform)
	} else {
		iter := client.Networking.Connect.Connections.ListAutoPaging(ctx, params, options...)
		return streamOutput("networking:connect:connections list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "networking:connect:connections list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleNetworkingConnectConnectionsDelete(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Networking.Connect.Connections.Delete(ctx, cmd.Value("connection-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:connect:connections delete", obj, format, transform)
}

func handleNetworkingConnectConnectionsGet(ctx context.Context, cmd *cli.Command) error {
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
		EmptyBody,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Networking.Connect.Connections.Get(ctx, cmd.Value("connection-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:connect:connections get", obj, format, transform)
}
