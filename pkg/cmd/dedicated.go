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

var rpcNodesDedicatedList = cli.Command{
	Name:  "list",
	Usage: "List all RPC Node Dedicated you created",
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
	Action:          handleRPCNodesDedicatedList,
	HideHelpCommand: true,
}

var rpcNodesDedicatedGet = cli.Command{
	Name:  "get",
	Usage: "Get details about an RPC Node Dedicated",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "node-id",
		},
	},
	Action:          handleRPCNodesDedicatedGet,
	HideHelpCommand: true,
}

func handleRPCNodesDedicatedList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := rpc_nodes.DedicatedListParams{}

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
		_, err = client.RPCNodes.Dedicated.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "rpc-nodes:dedicated list", obj, format, transform)
	} else {
		iter := client.RPCNodes.Dedicated.ListAutoPaging(ctx, params, options...)
		return streamOutput("rpc-nodes:dedicated list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "rpc-nodes:dedicated list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleRPCNodesDedicatedGet(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.RPCNodes.Dedicated.Get(ctx, cmd.Value("node-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "rpc-nodes:dedicated get", obj, format, transform)
}
