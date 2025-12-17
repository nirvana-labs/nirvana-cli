// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/api_keys"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var apiKeysCreate = cli.Command{
	Name:  "create",
	Usage: "Create a new API key",
	Flags: []cli.Flag{
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:     "expires-at",
			Usage:    "When the API Key expires and is no longer valid.",
			BodyPath: "expires_at",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "API Key name.",
			BodyPath: "name",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:     "starts-at",
			Usage:    "When the API Key starts to be valid.",
			BodyPath: "starts_at",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the API Key.",
			BodyPath: "tags",
		},
	},
	Action:          handleAPIKeysCreate,
	HideHelpCommand: true,
}

var apiKeysUpdate = cli.Command{
	Name:  "update",
	Usage: "Update an existing API key",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "api-key-id",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "API Key name.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the API Key.",
			BodyPath: "tags",
		},
	},
	Action:          handleAPIKeysUpdate,
	HideHelpCommand: true,
}

var apiKeysList = cli.Command{
	Name:  "list",
	Usage: "List all API keys for the authenticated user",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor returned by a previous request",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
	},
	Action:          handleAPIKeysList,
	HideHelpCommand: true,
}

var apiKeysDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete an API key",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "api-key-id",
		},
	},
	Action:          handleAPIKeysDelete,
	HideHelpCommand: true,
}

var apiKeysGet = cli.Command{
	Name:  "get",
	Usage: "Get details about an API key",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "api-key-id",
		},
	},
	Action:          handleAPIKeysGet,
	HideHelpCommand: true,
}

func handleAPIKeysCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := api_keys.APIKeyNewParams{}

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
	_, err = client.APIKeys.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "api-keys create", obj, format, transform)
}

func handleAPIKeysUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("api-key-id") && len(unusedArgs) > 0 {
		cmd.Set("api-key-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := api_keys.APIKeyUpdateParams{}

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
	_, err = client.APIKeys.Update(
		ctx,
		cmd.Value("api-key-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "api-keys update", obj, format, transform)
}

func handleAPIKeysList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := api_keys.APIKeyListParams{}

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

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.APIKeys.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "api-keys list", obj, format, transform)
	} else {
		iter := client.APIKeys.ListAutoPaging(ctx, params, options...)
		return streamOutput("api-keys list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "api-keys list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleAPIKeysDelete(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("api-key-id") && len(unusedArgs) > 0 {
		cmd.Set("api-key-id", unusedArgs[0])
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

	return client.APIKeys.Delete(ctx, cmd.Value("api-key-id").(string), options...)
}

func handleAPIKeysGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("api-key-id") && len(unusedArgs) > 0 {
		cmd.Set("api-key-id", unusedArgs[0])
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
	_, err = client.APIKeys.Get(ctx, cmd.Value("api-key-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "api-keys get", obj, format, transform)
}
