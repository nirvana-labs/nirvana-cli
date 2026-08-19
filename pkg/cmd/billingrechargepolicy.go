// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/nirvana-labs/nirvana-go/organizations"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var organizationsBillingRechargePolicyUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update the organization's recharge mode: manual for self-serve top-ups, or\nautomatic to charge the card on file at the recharge threshold (fixed and\nproportional required).",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[string]{
			Name:     "policy",
			Usage:    "Policy is the top-up mode.",
			Required: true,
			BodyPath: "policy",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "policy-args",
			Usage:    "PolicyArgs carries the threshold parameters. Required when policy is\n\"automatic\"; must be omitted when policy is \"manual\".",
			BodyPath: "policy_args",
		},
	},
	Action:          handleOrganizationsBillingRechargePolicyUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"policy-args": {
		&requestflag.InnerFlag[string]{
			Name:       "policy-args.fixed",
			Usage:      `Arbitrary-precision decimal serialized as a string (e.g. "58.40").`,
			InnerField: "fixed",
		},
		&requestflag.InnerFlag[string]{
			Name:       "policy-args.runway-days",
			Usage:      `Arbitrary-precision decimal serialized as a string (e.g. "58.40").`,
			InnerField: "runway_days",
		},
		&requestflag.InnerFlag[*string]{
			Name:       "policy-args.monthly-cap",
			Usage:      `Arbitrary-precision decimal serialized as a string (e.g. "58.40").`,
			InnerField: "monthly_cap",
		},
	},
})

var organizationsBillingRechargePolicyGet = cli.Command{
	Name:    "get",
	Usage:   "Get the organization's recharge configuration: the top-up mode, the fixed and\nproportional threshold components, and when the current mode took effect.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
	},
	Action:          handleOrganizationsBillingRechargePolicyGet,
	HideHelpCommand: true,
}

func handleOrganizationsBillingRechargePolicyUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("organization-id") && len(unusedArgs) > 0 {
		cmd.Set("organization-id", unusedArgs[0])
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
		false,
	)
	if err != nil {
		return err
	}

	params := organizations.BillingRechargePolicyUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Billing.RechargePolicy.Update(
		ctx,
		cmd.Value("organization-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "organizations:billing:recharge-policy update",
		Transform:      transform,
	})
}

func handleOrganizationsBillingRechargePolicyGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("organization-id") && len(unusedArgs) > 0 {
		cmd.Set("organization-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Billing.RechargePolicy.Get(ctx, cmd.Value("organization-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "organizations:billing:recharge-policy get",
		Transform:      transform,
	})
}
