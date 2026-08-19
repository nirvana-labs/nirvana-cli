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

var organizationsBillingCost = cli.Command{
	Name:    "cost",
	Usage:   "Get the organization's total usage cost per UTC day over a date range (max 90\ndays), summing open and closed resources. One entry per day, oldest first.\nDefaults to the last 30 days.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[any]{
			Name:      "from",
			Usage:     "Inclusive start day, YYYY-MM-DD (UTC). Defaults to 30 days before to.",
			QueryPath: "from",
		},
		&requestflag.Flag[any]{
			Name:      "to",
			Usage:     "Inclusive end day, YYYY-MM-DD (UTC). Defaults to today.",
			QueryPath: "to",
		},
	},
	Action:          handleOrganizationsBillingCost,
	HideHelpCommand: true,
}

var organizationsBillingHistory = cli.Command{
	Name:    "history",
	Usage:   "List the organization's billing history: prepaid credits, top-ups, and manual\nadjustments, newest first. Paginated with an opaque cursor.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[any]{
			Name:      "created-at-max",
			Usage:     "Only entries at or before this RFC 3339 instant",
			QueryPath: "created_at_max",
		},
		&requestflag.Flag[any]{
			Name:      "created-at-min",
			Usage:     "Only entries at or after this RFC 3339 instant",
			QueryPath: "created_at_min",
		},
		&requestflag.Flag[string]{
			Name:      "currency",
			Usage:     "Filter by currency, as an ISO 4217 code. Case-insensitive.",
			QueryPath: "currency",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor returned by a previous request. Only valid for the same filters and sort order.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "purpose",
			Usage:     "Filter by the funding flow a credit came from",
			QueryPath: "purpose",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, amount",
			Default:   "created_at:desc",
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "type",
			Usage:     "Filter by entry type",
			QueryPath: "type",
		},
	},
	Action:          handleOrganizationsBillingHistory,
	HideHelpCommand: true,
}

var organizationsBillingRecharge = cli.Command{
	Name:    "recharge",
	Usage:   "Charge the card on file up to the recharge target now instead of waiting for the\nscheduled retry. Automatic-policy prepaid organizations only. Idempotency-Key\nheader required.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleOrganizationsBillingRecharge,
	HideHelpCommand: true,
}

var organizationsBillingStatements = cli.Command{
	Name:    "statements",
	Usage:   "Get the itemized monthly usage statement: consumption grouped by project,\nresource type, and dimension, priced from recorded usage. Defaults to the\ncurrent month.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[string]{
			Name:      "month",
			Usage:     "Billing month, YYYY-MM (UTC). Defaults to the current month.",
			QueryPath: "month",
		},
	},
	Action:          handleOrganizationsBillingStatements,
	HideHelpCommand: true,
}

var organizationsBillingSummary = cli.Command{
	Name:    "summary",
	Usage:   "Get the organization's billing summary: effective balance, monthly and daily\nrun-rate cost, runway, and the projected next-recharge date. Costs are run-rate\nprojections.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
	},
	Action:          handleOrganizationsBillingSummary,
	HideHelpCommand: true,
}

var organizationsBillingTopUp = cli.Command{
	Name:    "top-up",
	Usage:   "Charge the card on file and credit the prepaid balance. A unique Idempotency-Key\nheader is required; reuse it across retries so a timed-out top-up is not charged\ntwice.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[string]{
			Name:     "amount",
			Usage:    "Amount to charge and credit, in USD. Must be greater than 0, at most two decimal places, and at most 10000.",
			Required: true,
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:       "idempotency-key",
			Required:   true,
			HeaderPath: "Idempotency-Key",
		},
	},
	Action:          handleOrganizationsBillingTopUp,
	HideHelpCommand: true,
}

func handleOrganizationsBillingCost(ctx context.Context, cmd *cli.Command) error {
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

	params := organizations.BillingCostParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Billing.Cost(
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
		Title:          "organizations:billing cost",
		Transform:      transform,
	})
}

func handleOrganizationsBillingHistory(ctx context.Context, cmd *cli.Command) error {
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

	params := organizations.BillingHistoryParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Billing.History(
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
		Title:          "organizations:billing history",
		Transform:      transform,
	})
}

func handleOrganizationsBillingRecharge(ctx context.Context, cmd *cli.Command) error {
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

	params := organizations.BillingRechargeParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Billing.Recharge(
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
		Title:          "organizations:billing recharge",
		Transform:      transform,
	})
}

func handleOrganizationsBillingStatements(ctx context.Context, cmd *cli.Command) error {
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

	params := organizations.BillingStatementsParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Billing.Statements(
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
		Title:          "organizations:billing statements",
		Transform:      transform,
	})
}

func handleOrganizationsBillingSummary(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Organizations.Billing.Summary(ctx, cmd.Value("organization-id").(string), options...)
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
		Title:          "organizations:billing summary",
		Transform:      transform,
	})
}

func handleOrganizationsBillingTopUp(ctx context.Context, cmd *cli.Command) error {
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

	params := organizations.BillingTopUpParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Billing.TopUp(
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
		Title:          "organizations:billing top-up",
		Transform:      transform,
	})
}
