// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/regions"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var regionsList = cli.Command{
	Name:    "list",
	Usage:   "List all regions",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "availability",
			Usage:     "Filter by region availability",
			QueryPath: "availability",
		},
		&requestflag.Flag[bool]{
			Name:      "compute-vms",
			Usage:     "Only regions where Virtual Machines are available",
			QueryPath: "compute_vms",
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
		&requestflag.Flag[bool]{
			Name:      "networking-connect",
			Usage:     "Only regions where Nirvana Connect is available",
			QueryPath: "networking_connect",
		},
		&requestflag.Flag[bool]{
			Name:      "networking-vpcs",
			Usage:     "Only regions where VPCs are available",
			QueryPath: "networking_vpcs",
		},
		&requestflag.Flag[bool]{
			Name:      "nks-autoscaling",
			Usage:     "Only regions where NKS node pool autoscaling is available",
			QueryPath: "nks_autoscaling",
		},
		&requestflag.Flag[bool]{
			Name:      "nks-clusters",
			Usage:     "Only regions where NKS clusters are available",
			QueryPath: "nks_clusters",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: longitude, name, availability",
			Default:   "longitude:asc",
			QueryPath: "sort",
		},
		&requestflag.Flag[bool]{
			Name:      "storage-abs",
			Usage:     "Only regions where Accelerated Block Storage is available",
			QueryPath: "storage_abs",
		},
		&requestflag.Flag[bool]{
			Name:      "storage-local-nvme",
			Usage:     "Only regions where locally-attached NVMe storage is available",
			QueryPath: "storage_local_nvme",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleRegionsList,
	HideHelpCommand: true,
}

var regionsGet = cli.Command{
	Name:    "get",
	Usage:   "Get a region by name",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "name",
			Required:  true,
			PathParam: "name",
		},
	},
	Action:          handleRegionsGet,
	HideHelpCommand: true,
}

func handleRegionsList(ctx context.Context, cmd *cli.Command) error {
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

	params := regions.RegionListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Regions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "regions list",
			Transform:      transform,
		})
	} else {
		iter := client.Regions.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "regions list",
			Transform:      transform,
		})
	}
}

func handleRegionsGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("name") && len(unusedArgs) > 0 {
		cmd.Set("name", unusedArgs[0])
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
	_, err = client.Regions.Get(ctx, cmd.Value("name").(string), options...)
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
		Title:          "regions get",
		Transform:      transform,
	})
}
