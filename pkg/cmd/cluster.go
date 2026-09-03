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

var nksClustersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create an NKS Cluster",
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
	Action:          handleNKSClustersCreate,
	HideHelpCommand: true,
}

var nksClustersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update an NKS cluster",
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
	Action:          handleNKSClustersUpdate,
	HideHelpCommand: true,
}

var nksClustersList = cli.Command{
	Name:    "list",
	Usage:   "List all NKS clusters",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project ID of resources to request",
			Required:  true,
			QueryPath: "project_id",
		},
		&requestflag.Flag[bool]{
			Name:      "autoscaling",
			Usage:     "Filter by whether autoscaling is enabled",
			QueryPath: "autoscaling",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor returned by a previous request. Only valid for the same filters and sort order.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[string]{
			Name:      "kubernetes-version",
			Usage:     "Filter by Kubernetes version, matched exactly",
			QueryPath: "kubernetes_version",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Filter by a case-insensitive substring of the Cluster name",
			QueryPath: "name",
		},
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Filter by region",
			QueryPath: "region",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, updated_at, name, status",
			Default:   "created_at:desc",
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by Cluster status",
			QueryPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Filter by tags. Repeat the parameter to require several tags; a Cluster must carry all of them.",
			QueryPath: "tags",
		},
		&requestflag.Flag[string]{
			Name:      "vpc-id",
			Usage:     "Filter by the VPC the Cluster is in",
			QueryPath: "vpc_id",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleNKSClustersList,
	HideHelpCommand: true,
}

var nksClustersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an NKS cluster",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cluster-id",
			Required:  true,
			PathParam: "cluster_id",
		},
	},
	Action:          handleNKSClustersDelete,
	HideHelpCommand: true,
}

var nksClustersGet = cli.Command{
	Name:    "get",
	Usage:   "Get details about an NKS cluster",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cluster-id",
			Required:  true,
			PathParam: "cluster_id",
		},
	},
	Action:          handleNKSClustersGet,
	HideHelpCommand: true,
}

func handleNKSClustersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.New(ctx, params, options...)
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
		Title:          "nks:clusters create",
		Transform:      transform,
	})
}

func handleNKSClustersUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Update(
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
		Title:          "nks:clusters update",
		Transform:      transform,
	})
}

func handleNKSClustersList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := nks.ClusterListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.NKS.Clusters.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "nks:clusters list",
			Transform:      transform,
		})
	} else {
		iter := client.NKS.Clusters.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "nks:clusters list",
			Transform:      transform,
		})
	}
}

func handleNKSClustersDelete(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Delete(ctx, cmd.Value("cluster-id").(string), options...)
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
		Title:          "nks:clusters delete",
		Transform:      transform,
	})
}

func handleNKSClustersGet(ctx context.Context, cmd *cli.Command) error {
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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.NKS.Clusters.Get(ctx, cmd.Value("cluster-id").(string), options...)
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
		Title:          "nks:clusters get",
		Transform:      transform,
	})
}
