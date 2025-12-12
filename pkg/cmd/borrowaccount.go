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

var vektorBorrowAccountsList = cli.Command{
	Name:  "list",
	Usage: "Get account-level borrow info on specified accounts",
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "account",
			Usage:    "A list of accounts. Currently only EVM addresses are supported.",
			BodyPath: "accounts",
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
		&requestflag.Flag[string]{
			Name:     "quote-asset-symbol",
			Usage:    "An asset symbol",
			BodyPath: "quote_asset_symbol",
		},
	},
	Action:          handleVektorBorrowAccountsList,
	HideHelpCommand: true,
}

var vektorBorrowAccountsListHistorical = cli.Command{
	Name:  "list-historical",
	Usage: "Get account-level borrow info on specified accounts",
	Flags: []cli.Flag{
		&requestflag.Flag[[]string]{
			Name:     "account",
			Usage:    "A list of accounts. Currently only EVM addresses are supported.",
			BodyPath: "accounts",
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
		&requestflag.Flag[string]{
			Name:     "quote-asset-symbol",
			Usage:    "An asset symbol",
			BodyPath: "quote_asset_symbol",
		},
	},
	Action:          handleVektorBorrowAccountsListHistorical,
	HideHelpCommand: true,
}

func handleVektorBorrowAccountsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := vektor.BorrowAccountListParams{}

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
	_, err = client.Vektor.Borrow.Accounts.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:borrow:accounts list", obj, format, transform)
}

func handleVektorBorrowAccountsListHistorical(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := vektor.BorrowAccountListHistoricalParams{}

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
	_, err = client.Vektor.Borrow.Accounts.ListHistorical(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:borrow:accounts list-historical", obj, format, transform)
}
