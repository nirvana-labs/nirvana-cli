// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/rpc_nodes"
	"github.com/stainless-sdks/nirvana-cli/internal/apiquery"
	"github.com/stainless-sdks/nirvana-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var rpcNodesFlexCreate = cli.Command{
	Name:  "create",
	Usage: "Create a new RPC Node Flex",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "blockchain",
			Usage:    "Blockchain.",
			BodyPath: "blockchain",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the RPC Node Flex.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "network",
			Usage:    "Network type (e.g., mainnet, testnet).",
			BodyPath: "network",
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID to associate with the RPC Node Flex.",
			BodyPath: "project_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the RPC Node Flex (optional, max 50).",
			BodyPath: "tags",
		},
	},
	Action:          handleRPCNodesFlexCreate,
	HideHelpCommand: true,
}

var rpcNodesFlexUpdate = cli.Command{
	Name:  "update",
	Usage: "Update an existing RPC Node Flex",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "node-id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the RPC Node Flex.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the RPC Node Flex (optional, max 50).",
			BodyPath: "tags",
		},
	},
	Action:          handleRPCNodesFlexUpdate,
	HideHelpCommand: true,
}

var rpcNodesFlexList = cli.Command{
	Name:  "list",
	Usage: "List all RPC Node Flex you created",
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
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project ID of resources to request",
			QueryPath: "project_id",
		},
	},
	Action:          handleRPCNodesFlexList,
	HideHelpCommand: true,
}

var rpcNodesFlexDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete an RPC Node Flex",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "node-id",
		},
	},
	Action:          handleRPCNodesFlexDelete,
	HideHelpCommand: true,
}

var rpcNodesFlexGet = cli.Command{
	Name:  "get",
	Usage: "Get details about an RPC Node Flex",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "node-id",
		},
	},
	Action:          handleRPCNodesFlexGet,
	HideHelpCommand: true,
}

func handleRPCNodesFlexCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := rpc_nodes.FlexNewParams{}

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
	_, err = client.RPCNodes.Flex.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rpc-nodes:flex create", obj, format, transform)
}

func handleRPCNodesFlexUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("node-id") && len(unusedArgs) > 0 {
		cmd.Set("node-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := rpc_nodes.FlexUpdateParams{}

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
	_, err = client.RPCNodes.Flex.Update(
		ctx,
		cmd.Value("node-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rpc-nodes:flex update", obj, format, transform)
}

func handleRPCNodesFlexList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := rpc_nodes.FlexListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.RPCNodes.Flex.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "rpc-nodes:flex list", obj, format, transform)
	} else {
		iter := client.RPCNodes.Flex.ListAutoPaging(ctx, params, options...)
		return streamOutput("rpc-nodes:flex list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "rpc-nodes:flex list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleRPCNodesFlexDelete(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("node-id") && len(unusedArgs) > 0 {
		cmd.Set("node-id", unusedArgs[0])
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
	)
	if err != nil {
		return err
	}

	return client.RPCNodes.Flex.Delete(ctx, cmd.Value("node-id").(string), options...)
}

func handleRPCNodesFlexGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("node-id") && len(unusedArgs) > 0 {
		cmd.Set("node-id", unusedArgs[0])
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
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.RPCNodes.Flex.Get(ctx, cmd.Value("node-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rpc-nodes:flex get", obj, format, transform)
}
