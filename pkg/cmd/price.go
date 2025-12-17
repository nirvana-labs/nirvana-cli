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

var vektorPricesList = cli.Command{
	Name:  "list",
	Usage: "Get a list of asset prices",
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "asset-symbol",
			Usage:    "A list of asset symbols",
			BodyPath: "asset_symbols",
		},
		&requestflag.Flag[string]{
			Name:     "quote-asset-symbol",
			Usage:    "An asset symbol",
			BodyPath: "quote_asset_symbol",
		},
	},
	Action:          handleVektorPricesList,
	HideHelpCommand: true,
}

var vektorPricesListHistorical = cli.Command{
	Name:  "list-historical",
	Usage: "Get a list of asset prices",
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "asset-symbol",
			Usage:    "A list of asset symbols",
			BodyPath: "asset_symbols",
		},
		&requestflag.Flag[string]{
			Name:     "quote-asset-symbol",
			Usage:    "An asset symbol",
			BodyPath: "quote_asset_symbol",
		},
	},
	Action:          handleVektorPricesListHistorical,
	HideHelpCommand: true,
}

func handleVektorPricesList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vektor.PriceListParams{}

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
	_, err = client.Vektor.Prices.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:prices list", obj, format, transform)
}

func handleVektorPricesListHistorical(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := vektor.PriceListHistoricalParams{}

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
	_, err = client.Vektor.Prices.ListHistorical(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:prices list-historical", obj, format, transform)
}
