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

var vektorSellSellCreate = cli.Command{
	Name:  "create",
	Usage: "Sell an asset for another asset",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "blockchain",
			Usage:    "A blockchain ID, represented as a TypeID with `blockchain` prefix",
			BodyPath: "blockchain",
		},
		&requestflag.Flag[string]{
			Name:     "from",
			Usage:    "An EVM address",
			BodyPath: "from",
		},
		&requestflag.Flag[string]{
			Name:     "receive-asset",
			Usage:    "An asset ID, represented as a TypeID with `asset` prefix",
			BodyPath: "receive_asset",
		},
		&requestflag.Flag[string]{
			Name:     "spend-amount",
			Usage:    "An arbitrary precision decimal represented as a string",
			BodyPath: "spend_amount",
		},
		&requestflag.Flag[string]{
			Name:     "spend-asset",
			Usage:    "An asset ID, represented as a TypeID with `asset` prefix",
			BodyPath: "spend_asset",
		},
		&requestflag.Flag[[]string]{
			Name:     "venue",
			Usage:    "A list of venue IDs or venue symbols",
			BodyPath: "venues",
		},
		&requestflag.Flag[string]{
			Name:     "slippage",
			Usage:    "An arbitrary precision decimal represented as a string",
			Default:  "0.01",
			BodyPath: "slippage",
		},
	},
	Action:          handleVektorSellSellCreate,
	HideHelpCommand: true,
}

func handleVektorSellSellCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vektor.SellSellNewParams{}

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

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Vektor.Sell.Sell.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:sell:sell create", obj, format, transform)
}
