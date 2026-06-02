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

var organizationsAddressCreate = cli.Command{
	Name:    "create",
	Usage:   "Create the address for an organization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[string]{
			Name:     "city",
			Usage:    "City or locality.",
			Required: true,
			BodyPath: "city",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Two-letter ISO 3166-1 alpha-2 country code.",
			Required: true,
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "line1",
			Usage:    "First line of the street address.",
			Required: true,
			BodyPath: "line1",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Usage:    "Postal or ZIP code.",
			Required: true,
			BodyPath: "postal_code",
		},
		&requestflag.Flag[string]{
			Name:     "line2",
			Usage:    "Second line of the street address (suite, unit, building).",
			BodyPath: "line2",
		},
		&requestflag.Flag[string]{
			Name:     "state",
			Usage:    "State, province, or region. Required by some tax jurisdictions (e.g. US, CA).",
			BodyPath: "state",
		},
		&requestflag.Flag[string]{
			Name:     "tax-id",
			Usage:    "Tax identification number (e.g. VAT, EIN, ABN). Optional.",
			BodyPath: "tax_id",
		},
		&requestflag.Flag[string]{
			Name:     "tax-id-type",
			Usage:    "Type of the tax identification number (e.g. eu_vat, us_ein, gb_vat, au_abn). Optional.",
			BodyPath: "tax_id_type",
		},
	},
	Action:          handleOrganizationsAddressCreate,
	HideHelpCommand: true,
}

var organizationsAddressUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update the address for an organization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[string]{
			Name:     "city",
			Usage:    "City or locality.",
			BodyPath: "city",
		},
		&requestflag.Flag[string]{
			Name:     "country",
			Usage:    "Two-letter ISO 3166-1 alpha-2 country code.",
			BodyPath: "country",
		},
		&requestflag.Flag[string]{
			Name:     "line1",
			Usage:    "First line of the street address.",
			BodyPath: "line1",
		},
		&requestflag.Flag[*string]{
			Name:     "line2",
			Usage:    "Second line of the street address (suite, unit, building). Omit to leave\nunchanged, send null to clear, or send a value to set it.",
			BodyPath: "line2",
		},
		&requestflag.Flag[string]{
			Name:     "postal-code",
			Usage:    "Postal or ZIP code.",
			BodyPath: "postal_code",
		},
		&requestflag.Flag[*string]{
			Name:     "state",
			Usage:    "State, province, or region. Omit to leave unchanged, send null to clear,\nor send a value to set it.",
			BodyPath: "state",
		},
		&requestflag.Flag[*string]{
			Name:     "tax-id",
			Usage:    "Tax identification number (e.g. VAT, EIN, ABN). Omit to leave unchanged,\nsend null to clear, or send a value to set it.",
			BodyPath: "tax_id",
		},
		&requestflag.Flag[*string]{
			Name:     "tax-id-type",
			Usage:    "Type of the tax identification number (e.g. eu_vat, us_ein, gb_vat, au_abn).\nOmit to leave unchanged, send null to clear, or send a value to set it.",
			BodyPath: "tax_id_type",
		},
	},
	Action:          handleOrganizationsAddressUpdate,
	HideHelpCommand: true,
}

var organizationsAddressGet = cli.Command{
	Name:    "get",
	Usage:   "Get the address for an organization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
	},
	Action:          handleOrganizationsAddressGet,
	HideHelpCommand: true,
}

func handleOrganizationsAddressCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := organizations.AddressNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Address.New(
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
		Title:          "organizations:address create",
		Transform:      transform,
	})
}

func handleOrganizationsAddressUpdate(ctx context.Context, cmd *cli.Command) error {
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

	params := organizations.AddressUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Address.Update(
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
		Title:          "organizations:address update",
		Transform:      transform,
	})
}

func handleOrganizationsAddressGet(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Organizations.Address.Get(ctx, cmd.Value("organization-id").(string), options...)
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
		Title:          "organizations:address get",
		Transform:      transform,
	})
}
