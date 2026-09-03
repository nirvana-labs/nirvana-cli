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

var nksClustersPoolsVolumesList = cli.Command{
	Name:    "list",
	Usage:   "List all volumes attached to the nodes of an NKS node pool",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cluster-id",
			Required:  true,
			PathParam: "cluster_id",
		},
		&requestflag.Flag[string]{
			Name:      "pool-id",
			Required:  true,
			PathParam: "pool_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor returned by a previous request. Only valid for the same filters and sort order.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "kind",
			Usage:     "Filter by volume kind",
			QueryPath: "kind",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Filter by a case-insensitive substring of the volume name",
			QueryPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:      "size-gb-max",
			Usage:     "Only volumes of at most this size",
			QueryPath: "size_gb_max",
		},
		&requestflag.Flag[int64]{
			Name:      "size-gb-min",
			Usage:     "Only volumes of at least this size",
			QueryPath: "size_gb_min",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, updated_at, name, status, size_gb",
			Default:   "created_at:desc",
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by volume status",
			QueryPath: "status",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Filter by storage type",
			QueryPath: "type",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleNKSClustersPoolsVolumesList,
	HideHelpCommand: true,
}

func handleNKSClustersPoolsVolumesList(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterPoolVolumeListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.NKS.Clusters.Pools.Volumes.List(
			ctx,
			cmd.Value("cluster-id").(string),
			cmd.Value("pool-id").(string),
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
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "nks:clusters:pools:volumes list",
			Transform:      transform,
		})
	} else {
		iter := client.NKS.Clusters.Pools.Volumes.ListAutoPaging(
			ctx,
			cmd.Value("cluster-id").(string),
			cmd.Value("pool-id").(string),
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
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "nks:clusters:pools:volumes list",
			Transform:      transform,
		})
	}
}
