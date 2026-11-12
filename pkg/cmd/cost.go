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

var sandboxClustersCostCreate = cli.Command{
	Name:    "create",
	Usage:   "Return a priced cost quote for the proposed Sandbox Cluster.",
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
	Action:          handleSandboxClustersCostCreate,
	HideHelpCommand: true,
}

func handleSandboxClustersCostCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := sandbox_clusters.CostNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.SandboxClusters.Cost.New(ctx, params, options...)
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
		Title:          "sandbox-clusters:cost create",
		Transform:      transform,
	})
}
