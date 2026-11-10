// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go/v3"
	"github.com/nirvana-labs/nirvana-go/v3/storage"
	"github.com/urfave/cli/v3"
)

var storageBucketsAvailabilityCreate = cli.Command{
	Name:    "create",
	Usage:   "Check if a Bucket can be created",
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
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Bucket.",
			BodyPath: "tags",
		},
	},
	Action:          handleStorageBucketsAvailabilityCreate,
	HideHelpCommand: true,
}

func handleStorageBucketsAvailabilityCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := storage.BucketAvailabilityNewParams{}

	return client.Storage.Buckets.Availability.New(ctx, params, options...)
}
