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

var vektorExecutionsStepsGet = cli.Command{
	Name:  "get",
	Usage: "Get a step of an execution",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "execution-id",
		},
		&requestflag.Flag[string]{
			Name: "step-id",
		},
	},
	Action:          handleVektorExecutionsStepsGet,
	HideHelpCommand: true,
}

var vektorExecutionsStepsSign = cli.Command{
	Name:  "sign",
	Usage: "Sign an EVM transaction step",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "execution-id",
		},
		&requestflag.Flag[string]{
			Name: "step-id",
		},
		&requestflag.Flag[string]{
			Name:     "signed-payload",
			Usage:    "A hex string starting with 0x",
			BodyPath: "signed_payload",
		},
	},
	Action:          handleVektorExecutionsStepsSign,
	HideHelpCommand: true,
}

func handleVektorExecutionsStepsGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("execution-id") && len(unusedArgs) > 0 {
		cmd.Set("execution-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("step-id") && len(unusedArgs) > 0 {
		cmd.Set("step-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
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
	_, err = client.Vektor.Executions.Steps.Get(
		ctx,
		cmd.Value("execution-id").(string),
		cmd.Value("step-id").(string),
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "vektor:executions:steps get", obj, format, transform)
}

func handleVektorExecutionsStepsSign(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("execution-id") && len(unusedArgs) > 0 {
		cmd.Set("execution-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("step-id") && len(unusedArgs) > 0 {
		cmd.Set("step-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := vektor.ExecutionStepSignParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	return client.Vektor.Executions.Steps.Sign(
		ctx,
		cmd.Value("execution-id").(string),
		cmd.Value("step-id").(string),
		params,
		options...,
	)
}
