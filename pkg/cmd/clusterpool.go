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

var nksClustersPoolsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a node pool in an NKS cluster",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cluster-id",
			Required:  true,
			PathParam: "cluster_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the node pool.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "node-config",
			Usage:    "Node configuration.",
			Required: true,
			BodyPath: "node_config",
		},
		&requestflag.Flag[int64]{
			Name:     "node-count",
			Usage:    "Number of nodes. Must be between 0 and 100.",
			BodyPath: "node_count",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the node pool.",
			BodyPath: "tags",
		},
	},
	Action:          handleNKSClustersPoolsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"node-config": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "node-config.boot-volume",
			Usage:      "Boot volume configuration.",
			InnerField: "boot_volume",
		},
		&requestflag.InnerFlag[string]{
			Name:       "node-config.instance-type",
			Usage:      "Instance type name used for worker nodes. Immutable after pool creation.",
			InnerField: "instance_type",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "node-config.labels",
			Usage:      "Kubernetes labels to apply to each node in the pool. Each entry is \"key=value\".\nKeys under kubernetes.io, k8s.io, and nirvanalabs.io prefixes are reserved.\nImmutable after pool creation.",
			InnerField: "labels",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "node-config.taints",
			Usage:      "Kubernetes taints to apply to each node in the pool at creation time.\nEach entry is \"key=value:Effect\" where Effect is NoSchedule, PreferNoSchedule, or NoExecute.\nImmutable after pool creation.",
			InnerField: "taints",
		},
	},
})

var nksClustersPoolsUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an NKS node pool",
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
			Name:     "name",
			Usage:    "Name of the node pool.",
			BodyPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:     "node-count",
			Usage:    "Number of nodes.",
			BodyPath: "node_count",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the node pool.",
			BodyPath: "tags",
		},
	},
	Action:          handleNKSClustersPoolsUpdate,
	HideHelpCommand: true,
}

var nksClustersPoolsList = cli.Command{
	Name:    "list",
	Usage:   "List all node pools in an NKS cluster",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cluster-id",
			Required:  true,
			PathParam: "cluster_id",
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
	Action:          handleNKSClustersPoolsList,
	HideHelpCommand: true,
}

var nksClustersPoolsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an NKS node pool",
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
	},
	Action:          handleNKSClustersPoolsDelete,
	HideHelpCommand: true,
}

var nksClustersPoolsGet = cli.Command{
	Name:    "get",
	Usage:   "Get details about an NKS node pool",
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
	},
	Action:          handleNKSClustersPoolsGet,
	HideHelpCommand: true,
}

func handleNKSClustersPoolsCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cluster-id") && len(unusedArgs) > 0 {
		cmd.Set("cluster-id", unusedArgs[0])
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

	params := nks.ClusterPoolNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Pools.New(
		ctx,
		cmd.Value("cluster-id").(string),
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
		Title:          "nks:clusters:pools create",
		Transform:      transform,
	})
}

func handleNKSClustersPoolsUpdate(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	params := nks.ClusterPoolUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Pools.Update(
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
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "nks:clusters:pools update",
		Transform:      transform,
	})
}

func handleNKSClustersPoolsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cluster-id") && len(unusedArgs) > 0 {
		cmd.Set("cluster-id", unusedArgs[0])
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

	params := nks.ClusterPoolListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.NKS.Clusters.Pools.List(
			ctx,
			cmd.Value("cluster-id").(string),
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
			Title:          "nks:clusters:pools list",
			Transform:      transform,
		})
	} else {
		iter := client.NKS.Clusters.Pools.ListAutoPaging(
			ctx,
			cmd.Value("cluster-id").(string),
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
			Title:          "nks:clusters:pools list",
			Transform:      transform,
		})
	}
}

func handleNKSClustersPoolsDelete(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Pools.Delete(
		ctx,
		cmd.Value("cluster-id").(string),
		cmd.Value("pool-id").(string),
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
		Title:          "nks:clusters:pools delete",
		Transform:      transform,
	})
}

func handleNKSClustersPoolsGet(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Pools.Get(
		ctx,
		cmd.Value("cluster-id").(string),
		cmd.Value("pool-id").(string),
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
		Title:          "nks:clusters:pools get",
		Transform:      transform,
	})
}
