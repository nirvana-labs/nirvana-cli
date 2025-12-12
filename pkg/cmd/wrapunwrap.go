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

var vektorWrapUnwrapCreate = cli.Command{
	Name:  "create",
	Usage: "Unwrap the wrapped native asset",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "amount",
			Usage:    "An arbitrary precision decimal represented as a string",
			BodyPath: "amount",
		},
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
	},
	Action:          handleVektorWrapUnwrapCreate,
	HideHelpCommand: true,
}

func handleVektorWrapUnwrapCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := vektor.WrapUnwrapNewParams{}

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
	_, err = client.Vektor.Wrap.Unwrap.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:wrap:unwrap create", obj, format, transform)
}
