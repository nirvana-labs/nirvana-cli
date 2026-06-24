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

var nksClustersPoolsCostCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Return a priced cost quote for the proposed NKS node pool.",
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
	Action:          handleNKSClustersPoolsCostCreate,
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
			Usage:      "Instance type name used for worker nodes.",
			InnerField: "instance_type",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "node-config.labels",
			Usage:      "Kubernetes labels to apply to each node in the pool. Each entry is \"key=value\".\nKeys under kubernetes.io, k8s.io, and nirvanalabs.io prefixes are reserved.",
			InnerField: "labels",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "node-config.taints",
			Usage:      "Kubernetes taints to apply to each node in the pool at creation time.\nEach entry is \"key=value:Effect\" where Effect is NoSchedule, PreferNoSchedule, or NoExecute.\nTaints are immutable after pool creation.",
			InnerField: "taints",
		},
	},
})

var nksClustersPoolsCostUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Return a priced cost quote for the proposed NKS node pool update plus a diff\nagainst the current state.",
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
		&requestflag.Flag[map[string]any]{
			Name:     "node-config",
			Usage:    "Partial node configuration update.",
			BodyPath: "node_config",
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
	Action:          handleNKSClustersPoolsCostUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"node-config": {
		&requestflag.InnerFlag[[]string]{
			Name:       "node-config.labels",
			Usage:      "Kubernetes labels to apply to each node in the pool. Each entry is \"key=value\".\nWhen provided, the list fully replaces the current labels on the pool and on live nodes.",
			InnerField: "labels",
		},
	},
})

func handleNKSClustersPoolsCostCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterPoolCostNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Pools.Cost.New(
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
		Title:          "nks:clusters:pools:cost create",
		Transform:      transform,
	})
}

func handleNKSClustersPoolsCostUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterPoolCostUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Pools.Cost.Update(
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
		Title:          "nks:clusters:pools:cost update",
		Transform:      transform,
	})
}
