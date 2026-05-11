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

var nksClustersAvailabilityCreate = cli.Command{
	Name:    "create",
	Usage:   "Check if an NKS cluster can be created",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "autoscaling",
			Usage:    "Whether to enable autoscaling for the Cluster.",
			Required: true,
			BodyPath: "autoscaling",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Cluster.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID to create the Cluster in.",
			Required: true,
			BodyPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region the resource is in.",
			Required: true,
			BodyPath: "region",
		},
		&requestflag.Flag[string]{
			Name:     "vpc-id",
			Usage:    "ID of the VPC to use for the Cluster.",
			Required: true,
			BodyPath: "vpc_id",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Cluster.",
			BodyPath: "tags",
		},
	},
	Action:          handleNKSClustersAvailabilityCreate,
	HideHelpCommand: true,
}

var nksClustersAvailabilityUpdate = cli.Command{
	Name:    "update",
	Usage:   "Check if an NKS cluster can be updated",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cluster-id",
			Required:  true,
			PathParam: "cluster_id",
		},
		&requestflag.Flag[bool]{
			Name:     "autoscaling",
			Usage:    "Whether to enable autoscaling for the Cluster.",
			BodyPath: "autoscaling",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Cluster.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Cluster.",
			BodyPath: "tags",
		},
	},
	Action:          handleNKSClustersAvailabilityUpdate,
	HideHelpCommand: true,
}

func handleNKSClustersAvailabilityCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterAvailabilityNewParams{}

	return client.NKS.Clusters.Availability.New(ctx, params, options...)
}

func handleNKSClustersAvailabilityUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterAvailabilityUpdateParams{}

	return client.NKS.Clusters.Availability.Update(
		ctx,
		cmd.Value("cluster-id").(string),
		params,
		options...,
	)
}
