// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go/v3"
	"github.com/nirvana-labs/nirvana-go/v3/option"
	"github.com/nirvana-labs/nirvana-go/v3/sandbox_clusters"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var sandboxClustersCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a Sandbox Cluster to run sandboxes on. The infrastructure it runs on is\nprovisioned in a project the platform manages.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "instance-type",
			Usage:    "Instance type of the nodes the sandboxes run on.",
			Required: true,
			BodyPath: "instance_type",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Sandbox Cluster. Unique within its project.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:     "node-count",
			Usage:    "Number of nodes the sandboxes run on.",
			Required: true,
			BodyPath: "node_count",
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "ID of the project to create the Sandbox Cluster in.",
			Required: true,
			BodyPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region the resource is in.",
			Required: true,
			BodyPath: "region",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Sandbox Cluster.",
			BodyPath: "tags",
		},
	},
	Action:          handleSandboxClustersCreate,
	HideHelpCommand: true,
}

var sandboxClustersUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a Sandbox Cluster's name and tags",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sandbox-cluster-id",
			Required:  true,
			PathParam: "sandbox_cluster_id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Sandbox Cluster. Unique within its project.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Sandbox Cluster.",
			BodyPath: "tags",
		},
	},
	Action:          handleSandboxClustersUpdate,
	HideHelpCommand: true,
}

var sandboxClustersList = cli.Command{
	Name:    "list",
	Usage:   "List the Sandbox Clusters in a project",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "project-id",
			Usage:     "Project ID of resources to request",
			Required:  true,
			QueryPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor returned by a previous request. Only valid for the same sort order.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Filter by a case-insensitive substring of the Sandbox Cluster name",
			QueryPath: "name",
		},
		&requestflag.Flag[string]{
			Name:      "region",
			Usage:     "Filter by region",
			QueryPath: "region",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, updated_at",
			Default:   "created_at:desc",
			QueryPath: "sort",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Filter by tags. Repeat the parameter to require several tags; a Sandbox Cluster must carry all of them.",
			QueryPath: "tags",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleSandboxClustersList,
	HideHelpCommand: true,
}

var sandboxClustersDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a Sandbox Cluster, with its sandboxes and the infrastructure they run on.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sandbox-cluster-id",
			Required:  true,
			PathParam: "sandbox_cluster_id",
		},
	},
	Action:          handleSandboxClustersDelete,
	HideHelpCommand: true,
}

var sandboxClustersGet = cli.Command{
	Name:    "get",
	Usage:   "Get a Sandbox Cluster",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "sandbox-cluster-id",
			Required:  true,
			PathParam: "sandbox_cluster_id",
		},
	},
	Action:          handleSandboxClustersGet,
	HideHelpCommand: true,
}

func handleSandboxClustersCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := sandbox_clusters.SandboxClusterNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SandboxClusters.New(ctx, params, options...)
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
		Title:          "sandbox-clusters create",
		Transform:      transform,
	})
}

func handleSandboxClustersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sandbox-cluster-id") && len(unusedArgs) > 0 {
		cmd.Set("sandbox-cluster-id", unusedArgs[0])
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

	params := sandbox_clusters.SandboxClusterUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SandboxClusters.Update(
		ctx,
		cmd.Value("sandbox-cluster-id").(string),
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
		Title:          "sandbox-clusters update",
		Transform:      transform,
	})
}

func handleSandboxClustersList(ctx context.Context, cmd *cli.Command) error {
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

	params := sandbox_clusters.SandboxClusterListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.SandboxClusters.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "sandbox-clusters list",
			Transform:      transform,
		})
	} else {
		iter := client.SandboxClusters.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "sandbox-clusters list",
			Transform:      transform,
		})
	}
}

func handleSandboxClustersDelete(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sandbox-cluster-id") && len(unusedArgs) > 0 {
		cmd.Set("sandbox-cluster-id", unusedArgs[0])
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
	_, err = client.SandboxClusters.Delete(ctx, cmd.Value("sandbox-cluster-id").(string), options...)
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
		Title:          "sandbox-clusters delete",
		Transform:      transform,
	})
}

func handleSandboxClustersGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("sandbox-cluster-id") && len(unusedArgs) > 0 {
		cmd.Set("sandbox-cluster-id", unusedArgs[0])
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
	_, err = client.SandboxClusters.Get(ctx, cmd.Value("sandbox-cluster-id").(string), options...)
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
		Title:          "sandbox-clusters get",
		Transform:      transform,
	})
}
