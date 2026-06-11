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

var nksClustersCostCreate = cli.Command{
	Name:    "create",
	Usage:   "Return a priced cost quote for the proposed NKS cluster.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:     "autoscaling",
			Usage:    "Whether to enable autoscaling for the Cluster.",
			Required: true,
			BodyPath: "autoscaling",
		},
		&requestflag.Flag[string]{
			Name:     "kubernetes-version",
			Usage:    "Kubernetes version for the Cluster.",
			Required: true,
			BodyPath: "kubernetes_version",
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
	Action:          handleNKSClustersCostCreate,
	HideHelpCommand: true,
}

var nksClustersCostUpdate = cli.Command{
	Name:    "update",
	Usage:   "Return a priced cost quote for the proposed NKS cluster update plus a diff\nagainst the current state.",
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
	Action:          handleNKSClustersCostUpdate,
	HideHelpCommand: true,
}

func handleNKSClustersCostCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterCostNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Cost.New(ctx, params, options...)
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
		Title:          "nks:clusters:cost create",
		Transform:      transform,
	})
}

func handleNKSClustersCostUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterCostUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Cost.Update(
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
		Title:          "nks:clusters:cost update",
		Transform:      transform,
	})
}
