// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/vektor"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var vektorRegistryLPPoolsList = cli.Command{
	Name:  "list",
	Usage: "List LP pools in the registry, optionally filtered by blockchain, assets or\nvenues",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "A LP pool ID, represented as a TypeID with `lp_pool` prefix",
			BodyPath: "id",
		},
		&requestflag.Flag[[]string]{
			Name:     "asset",
			Usage:    "A list of asset IDs, EVM addresses or asset symbols",
			BodyPath: "assets",
		},
		&requestflag.Flag[string]{
			Name:     "blockchain",
			Usage:    "A blockchain ID, represented as a TypeID with `blockchain` prefix",
			BodyPath: "blockchain",
		},
		&requestflag.Flag[[]string]{
			Name:     "venue",
			Usage:    "A list of venue IDs or venue symbols",
			BodyPath: "venues",
		},
	},
	Action:          handleVektorRegistryLPPoolsList,
	HideHelpCommand: true,
}

func handleVektorRegistryLPPoolsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := vektor.RegistryLPPoolListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vektor.Registry.LPPools.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:registry:lp-pools list", obj, format, transform)
}
