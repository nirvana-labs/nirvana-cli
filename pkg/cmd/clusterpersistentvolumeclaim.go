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

var nksClustersPersistentVolumeClaimsList = cli.Command{
	Name:    "list",
	Usage:   "List all persistent volume claims in an NKS cluster",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cluster-id",
			Required:  true,
			PathParam: "cluster_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor returned by a previous request. Only valid for the same filters and sort order.",
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
			Usage:     "Filter by a case-insensitive substring of the claim name",
			QueryPath: "name",
		},
		&requestflag.Flag[int64]{
			Name:      "size-gb-max",
			Usage:     "Only claims of at most this size",
			QueryPath: "size_gb_max",
		},
		&requestflag.Flag[int64]{
			Name:      "size-gb-min",
			Usage:     "Only claims of at least this size",
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
			Usage:     "Filter by persistent volume claim status",
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
	Action:          handleNKSClustersPersistentVolumeClaimsList,
	HideHelpCommand: true,
}

var nksClustersPersistentVolumeClaimsGet = cli.Command{
	Name:    "get",
	Usage:   "Get details about an NKS persistent volume claim",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cluster-id",
			Required:  true,
			PathParam: "cluster_id",
		},
		&requestflag.Flag[string]{
			Name:      "persistent-volume-claim-id",
			Required:  true,
			PathParam: "persistent_volume_claim_id",
		},
	},
	Action:          handleNKSClustersPersistentVolumeClaimsGet,
	HideHelpCommand: true,
}

func handleNKSClustersPersistentVolumeClaimsList(ctx context.Context, cmd *cli.Command) error {
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

	params := nks.ClusterPersistentVolumeClaimListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.NKS.Clusters.PersistentVolumeClaims.List(
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
			Title:          "nks:clusters:persistent-volume-claims list",
			Transform:      transform,
		})
	} else {
		iter := client.NKS.Clusters.PersistentVolumeClaims.ListAutoPaging(
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
			Title:          "nks:clusters:persistent-volume-claims list",
			Transform:      transform,
		})
	}
}

func handleNKSClustersPersistentVolumeClaimsGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("cluster-id") && len(unusedArgs) > 0 {
		cmd.Set("cluster-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("persistent-volume-claim-id") && len(unusedArgs) > 0 {
		cmd.Set("persistent-volume-claim-id", unusedArgs[0])
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
	_, err = client.NKS.Clusters.PersistentVolumeClaims.Get(
		ctx,
		cmd.Value("cluster-id").(string),
		cmd.Value("persistent-volume-claim-id").(string),
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
		Title:          "nks:clusters:persistent-volume-claims get",
		Transform:      transform,
	})
}
