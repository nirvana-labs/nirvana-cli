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

var vektorBridgeQuotesList = cli.Command{
	Name:  "list",
	Usage: "Get quotes for bridging an exact amount of an asset to another blockchain",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "amount",
			Usage:    "An arbitrary precision decimal represented as a string",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "asset",
			Usage:    "An asset symbol",
			BodyPath: "asset",
		},
		&requestflag.Flag[string]{
			Name:     "from-blockchain",
			Usage:    "A blockchain ID, represented as a TypeID with `blockchain` prefix",
			BodyPath: "from_blockchain",
		},
		&requestflag.Flag[string]{
			Name:     "to-blockchain",
			Usage:    "A blockchain ID, represented as a TypeID with `blockchain` prefix",
			BodyPath: "to_blockchain",
		},
		&requestflag.Flag[[]string]{
			Name:     "venue",
			Usage:    "A list of venue IDs or venue symbols",
			BodyPath: "venues",
		},
		&requestflag.Flag[string]{
			Name:     "quote-asset-symbol",
			Usage:    "An asset symbol",
			BodyPath: "quote_asset_symbol",
		},
	},
	Action:          handleVektorBridgeQuotesList,
	HideHelpCommand: true,
}

func handleVektorBridgeQuotesList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := vektor.BridgeQuoteListParams{}

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
	_, err = client.Vektor.Bridge.Quotes.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:bridge:quotes list", obj, format, transform)
}
