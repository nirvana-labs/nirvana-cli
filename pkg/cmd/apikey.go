// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/api_keys"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var apiKeysCreate = requestflag.WithInnerFlags(cli.Command{
	Name:    "create",
	Usage:   "Create a new API key",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[any]{
			Name:     "expires-at",
			Usage:    "When the API Key expires and is no longer valid.",
			Required: true,
			BodyPath: "expires_at",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "API Key name.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "permission",
			Usage:    "Scoped permissions for this API key. At least one is required.",
			Required: true,
			BodyPath: "permissions",
		},
		&requestflag.Flag[[]string]{
			Name:     "project-id",
			Usage:    "Project IDs this API key is scoped to. At least one is required.",
			Required: true,
			BodyPath: "project_ids",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "source-ip-rule",
			Usage:    "IP filter rules.",
			BodyPath: "source_ip_rule",
		},
		&requestflag.Flag[any]{
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
}, map[string][]requestflag.HasOuterFlag{
	"permission": {
		&requestflag.InnerFlag[string]{
			Name:       "permission.permission",
			Usage:      `Permission level: "read" or "edit".`,
			InnerField: "permission",
		},
		&requestflag.InnerFlag[string]{
			Name:       "permission.resource-type",
			Usage:      "Resource type this permission applies to.",
			InnerField: "resource_type",
		},
	},
	"source-ip-rule": {
		&requestflag.InnerFlag[[]string]{
			Name:       "source-ip-rule.allowed",
			Usage:      "List of IPv4 CIDR addresses to allow.",
			InnerField: "allowed",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "source-ip-rule.blocked",
			Usage:      "List of IPv4 CIDR addresses to deny.",
			InnerField: "blocked",
		},
	},
})

var apiKeysUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:    "update",
	Usage:   "Update an existing API key",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "api-key-id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "API Key name.",
			BodyPath: "name",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "permission",
			Usage:    "Scoped permissions for this API key. When provided, replaces the entire set. At least one is required.",
			BodyPath: "permissions",
		},
		&requestflag.Flag[[]string]{
			Name:     "project-id",
			Usage:    "Project IDs this API key is scoped to. When provided, replaces the entire set. At least one is required.",
			BodyPath: "project_ids",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "source-ip-rule",
			Usage:    "IP filter rules.",
			BodyPath: "source_ip_rule",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the API Key.",
			BodyPath: "tags",
		},
	},
	Action:          handleAPIKeysUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"permission": {
		&requestflag.InnerFlag[string]{
			Name:       "permission.permission",
			Usage:      `Permission level: "read" or "edit".`,
			InnerField: "permission",
		},
		&requestflag.InnerFlag[string]{
			Name:       "permission.resource-type",
			Usage:      "Resource type this permission applies to.",
			InnerField: "resource_type",
		},
	},
	"source-ip-rule": {
		&requestflag.InnerFlag[[]string]{
			Name:       "source-ip-rule.allowed",
			Usage:      "List of IPv4 CIDR addresses to allow.",
			InnerField: "allowed",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "source-ip-rule.blocked",
			Usage:      "List of IPv4 CIDR addresses to deny.",
			InnerField: "blocked",
		},
	},
})

var apiKeysList = cli.Command{
	Name:    "list",
	Usage:   "List all API keys",
	Suggest: true,
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
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAPIKeysList,
	HideHelpCommand: true,
}

var apiKeysDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete an API key",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "api-key-id",
			Required: true,
		},
	},
	Action:          handleAPIKeysDelete,
	HideHelpCommand: true,
}

var apiKeysGet = cli.Command{
	Name:    "get",
	Usage:   "Get details about an API key",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "api-key-id",
			Required: true,
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "api-keys create",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "api-keys update",
		Transform:      transform,
	})
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.APIKeys.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "api-keys list",
			Transform:      transform,
		})
	} else {
		iter := client.APIKeys.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "api-keys list",
			Transform:      transform,
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
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(obj, ShowJSONOpts{
		ExplicitFormat: explicitFormat,
		Format:         format,
		RawOutput:      cmd.Root().Bool("raw-output"),
		Title:          "api-keys get",
		Transform:      transform,
	})
}
