// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/nks"
	"github.com/urfave/cli/v3"
)

var nksClustersPoolsAvailabilityCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Check if a node pool can be created in an NKS cluster",
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
	Action:          handleNKSClustersPoolsAvailabilityCreate,
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

var nksClustersPoolsAvailabilityUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Check if an NKS node pool can be updated",
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
	Action:          handleNKSClustersPoolsAvailabilityUpdate,
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

func handleNKSClustersPoolsAvailabilityCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterPoolAvailabilityNewParams{}

	return client.NKS.Clusters.Pools.Availability.New(
		ctx,
		cmd.Value("cluster-id").(string),
		params,
		options...,
	)
}

func handleNKSClustersPoolsAvailabilityUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterPoolAvailabilityUpdateParams{}

	return client.NKS.Clusters.Pools.Availability.Update(
		ctx,
		cmd.Value("cluster-id").(string),
		cmd.Value("pool-id").(string),
		params,
		options...,
	)
}
