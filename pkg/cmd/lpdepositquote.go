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

var vektorLPDepositQuoteCreate = cli.Command{
	Name:  "create",
	Usage: "Simulate depositing liquidity to a specific LP pool, creating a position or\nadding to an existing one.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "amount",
			Usage:    "An arbitrary precision decimal represented as a string",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "asset",
			Usage:    "An asset ID, represented as a TypeID with `asset` prefix",
			BodyPath: "asset",
		},
		&requestflag.Flag[string]{
			Name:     "blockchain",
			Usage:    "A blockchain ID, represented as a TypeID with `blockchain` prefix",
			BodyPath: "blockchain",
		},
		&requestflag.Flag[string]{
			Name:     "lp-pool-id",
			Usage:    "A LP pool ID, represented as a TypeID with `lp_pool` prefix",
			BodyPath: "lp_pool_id",
		},
		&requestflag.Flag[string]{
			Name:     "quote-asset-symbol",
			Usage:    "An asset symbol",
			BodyPath: "quote_asset_symbol",
		},
		&requestflag.Flag[any]{
			Name:     "range",
			Usage:    "A Uniswap V3 range. Lower and upper bounds should satisfy 0 <= `lower` < `upper`. The value -1 can be used in `upper` for infinity",
			BodyPath: "range",
		},
		&requestflag.Flag[string]{
			Name:     "account",
			Usage:    "An EVM address",
			BodyPath: "account",
		},
		&requestflag.Flag[any]{
			Name:     "specifier",
			Usage:    "Uniswap position specifier",
			BodyPath: "specifier",
		},
	},
	Action:          handleVektorLPDepositQuoteCreate,
	HideHelpCommand: true,
}

func handleVektorLPDepositQuoteCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := vektor.LPDepositQuoteNewParams{}

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
	_, err = client.Vektor.LP.DepositQuote.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:lp:deposit-quote create", obj, format, transform)
}
