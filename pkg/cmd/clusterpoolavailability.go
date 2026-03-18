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
			Name:     "cluster-id",
			Required: true,
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
			Usage:    "Number of nodes. Must be between 1 and 100.",
			Required: true,
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
		&requestflag.InnerFlag[int64]{
			Name:       "node-config.ram-gi",
			Usage:      "RAM size in GiB per node.",
			InnerField: "ram_gi",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "node-config.storage-gi",
			Usage:      "Storage size in GiB per node.",
			InnerField: "storage_gi",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "node-config.vcpu",
			Usage:      "Number of virtual CPUs per node.",
			InnerField: "vcpu",
		},
	},
})

var nksClustersPoolsAvailabilityUpdate = cli.Command{
	Name:    "update",
	Usage:   "Check if an NKS node pool can be updated",
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
			Name:     "name",
			Usage:    "Name of the node pool.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the node pool.",
			BodyPath: "tags",
		},
	},
	Action:          handleNKSClustersPoolsAvailabilityUpdate,
	HideHelpCommand: true,
}

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

	params := nks.ClusterPoolAvailabilityNewParams{}

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

	params := nks.ClusterPoolAvailabilityUpdateParams{}

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

	return client.NKS.Clusters.Pools.Availability.Update(
		ctx,
		cmd.Value("cluster-id").(string),
		cmd.Value("pool-id").(string),
		params,
		options...,
	)
}
