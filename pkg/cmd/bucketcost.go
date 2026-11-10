// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go/v3"
	"github.com/nirvana-labs/nirvana-go/v3/option"
	"github.com/nirvana-labs/nirvana-go/v3/storage"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var storageBucketsCostCreate = cli.Command{
	Name:    "create",
	Usage:   "Return a priced monthly cost quote for the proposed Bucket storing size_gb of\ndata. A Bucket is charged per GB it stores, prorated by the hour.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Bucket, used to address it over S3. It must be 3 to 63\ncharacters of lowercase letters, digits, dots and hyphens, start and end\nwith a letter or digit, not contain two adjacent dots and not be formatted\nas an IP address. \"admin\" and \"ops-api\" are reserved. The name cannot be\nchanged once the Bucket is created.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "project-id",
			Usage:    "Project ID the Bucket belongs to.",
			Required: true,
			BodyPath: "project_id",
		},
		&requestflag.Flag[string]{
			Name:     "region",
			Usage:    "Region the resource is in.",
			Required: true,
			BodyPath: "region",
		},
		&requestflag.Flag[int64]{
			Name:     "size-gb",
			Usage:    "Amount of data to price the Bucket at storing, in GB (1 GB = 10^9 bytes).\nA Bucket is charged for what it stores, so a quote without it is empty.",
			BodyPath: "size_gb",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Bucket.",
			BodyPath: "tags",
		},
	},
	Action:          handleStorageBucketsCostCreate,
	HideHelpCommand: true,
}

func handleStorageBucketsCostCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := storage.BucketCostNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Storage.Buckets.Cost.New(ctx, params, options...)
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
		Title:          "storage:buckets:cost create",
		Transform:      transform,
	})
}
