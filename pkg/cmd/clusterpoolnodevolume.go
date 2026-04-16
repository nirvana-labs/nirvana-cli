// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/nks"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var nksClustersPoolsNodesVolumesList = cli.Command{
	Name:    "list",
	Usage:   "List all volumes attached to an NKS node",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cluster-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "pool-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "node-id",
			Required: true,
		},
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleNKSClustersPoolsNodesVolumesList,
	HideHelpCommand: true,
}

var nksClustersPoolsNodesVolumesGet = cli.Command{
	Name:    "get",
	Usage:   "Get details about a volume attached to an NKS node",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "cluster-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "pool-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "node-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "volume-id",
			Required: true,
		},
	},
	Action:          handleNKSClustersPoolsNodesVolumesGet,
	HideHelpCommand: true,
}

func handleNKSClustersPoolsNodesVolumesList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cluster-id") && len(unusedArgs) > 0 {
		cmd.Set("cluster-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("pool-id") && len(unusedArgs) > 0 {
		cmd.Set("pool-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("node-id") && len(unusedArgs) > 0 {
		cmd.Set("node-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := nks.ClusterPoolNodeVolumeListParams{}

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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.NKS.Clusters.Pools.Nodes.Volumes.List(
			ctx,
			cmd.Value("cluster-id").(string),
			cmd.Value("pool-id").(string),
			cmd.Value("node-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "nks:clusters:pools:nodes:volumes list",
			Transform:      transform,
		})
	} else {
		iter := client.NKS.Clusters.Pools.Nodes.Volumes.ListAutoPaging(
			ctx,
			cmd.Value("cluster-id").(string),
			cmd.Value("pool-id").(string),
			cmd.Value("node-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			Title:          "nks:clusters:pools:nodes:volumes list",
			Transform:      transform,
		})
	}
}

func handleNKSClustersPoolsNodesVolumesGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cluster-id") && len(unusedArgs) > 0 {
		cmd.Set("cluster-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("pool-id") && len(unusedArgs) > 0 {
		cmd.Set("pool-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("node-id") && len(unusedArgs) > 0 {
		cmd.Set("node-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
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
	_, err = client.NKS.Clusters.Pools.Nodes.Volumes.Get(
		ctx,
		cmd.Value("cluster-id").(string),
		cmd.Value("pool-id").(string),
		cmd.Value("node-id").(string),
		cmd.Value("volume-id").(string),
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
		Title:          "nks:clusters:pools:nodes:volumes get",
		Transform:      transform,
	})
}
