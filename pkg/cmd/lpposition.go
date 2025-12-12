// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/vektor"
	"github.com/stainless-sdks/nirvana-cli/internal/apiquery"
	"github.com/stainless-sdks/nirvana-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var vektorLPPositionsList = cli.Command{
	Name:  "list",
	Usage: "Get info on LP positions",
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "account",
			Usage:    "A list of accounts. Currently only EVM addresses are supported.",
			BodyPath: "accounts",
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
		&requestflag.Flag[any]{
			Name:     "at",
			Usage:    "Either a ISO8601 timestamp or a block number",
			BodyPath: "at",
		},
		&requestflag.Flag[bool]{
			Name:     "exclude-zeros",
			Usage:    "Whether to exclude empty positions",
			BodyPath: "exclude_zeros",
		},
		&requestflag.Flag[string]{
			Name:     "quote-asset-symbol",
			Usage:    "An asset symbol",
			BodyPath: "quote_asset_symbol",
		},
		&requestflag.Flag[[]string]{
			Name:     "lp-pool-id",
			Usage:    "A list of LP pool IDs",
			BodyPath: "lp_pool_ids",
		},
	},
	Action:          handleVektorLPPositionsList,
	HideHelpCommand: true,
}

var vektorLPPositionsListHistorical = cli.Command{
	Name:  "list-historical",
	Usage: "Get info on LP positions",
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "account",
			Usage:    "A list of accounts. Currently only EVM addresses are supported.",
			BodyPath: "accounts",
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
		&requestflag.Flag[any]{
			Name:     "from",
			Usage:    "Either a ISO8601 timestamp or a block number",
			BodyPath: "from",
		},
		&requestflag.Flag[any]{
			Name:     "to",
			Usage:    "Either a ISO8601 timestamp or a block number",
			BodyPath: "to",
		},
		&requestflag.Flag[[]string]{
			Name:     "venue",
			Usage:    "A list of venue IDs or venue symbols",
			BodyPath: "venues",
		},
		&requestflag.Flag[bool]{
			Name:     "exclude-zeros",
			Usage:    "Whether to exclude empty positions",
			BodyPath: "exclude_zeros",
		},
		&requestflag.Flag[string]{
			Name:     "quote-asset-symbol",
			Usage:    "An asset symbol",
			BodyPath: "quote_asset_symbol",
		},
		&requestflag.Flag[[]string]{
			Name:     "lp-pool-id",
			Usage:    "A list of LP pool IDs",
			BodyPath: "lp_pool_ids",
		},
	},
	Action:          handleVektorLPPositionsListHistorical,
	HideHelpCommand: true,
}

func handleVektorLPPositionsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := vektor.LPPositionListParams{}

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
	_, err = client.Vektor.LP.Positions.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:lp:positions list", obj, format, transform)
}

func handleVektorLPPositionsListHistorical(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := vektor.LPPositionListHistoricalParams{}

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
	_, err = client.Vektor.LP.Positions.ListHistorical(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:lp:positions list-historical", obj, format, transform)
}
