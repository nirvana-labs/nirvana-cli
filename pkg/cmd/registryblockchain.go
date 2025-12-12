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

var vektorRegistryBlockchainsList = cli.Command{
	Name:  "list",
	Usage: "List supported blockchains",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Usage:    "A blockchain ID, represented as a TypeID with `blockchain` prefix",
			BodyPath: "id",
		},
		&requestflag.Flag[string]{
			Name:     "symbol",
			Usage:    "A blockchain symbol",
			BodyPath: "symbol",
		},
	},
	Action:          handleVektorRegistryBlockchainsList,
	HideHelpCommand: true,
}

func handleVektorRegistryBlockchainsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := vektor.RegistryBlockchainListParams{}

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
	_, err = client.Vektor.Registry.Blockchains.List(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:registry:blockchains list", obj, format, transform)
}
