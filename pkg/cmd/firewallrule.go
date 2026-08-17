// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var networkingFirewallRulesCreate = cli.Command{
	Name:    "create",
	Usage:   "Create a firewall rule",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "vpc-id",
			Required:  true,
			PathParam: "vpc_id",
		},
		&requestflag.Flag[string]{
			Name:     "destination-address",
			Usage:    "Destination address of the Firewall Rule. Either VPC CIDR or VM in VPC. Must be in network-aligned/canonical form.",
			Required: true,
			BodyPath: "destination_address",
		},
		&requestflag.Flag[[]string]{
			Name:     "destination-port",
			Usage:    "Destination ports of the Firewall Rule.",
			Required: true,
			BodyPath: "destination_ports",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Firewall Rule.",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "protocol",
			Usage:    "Protocol of the Firewall Rule.",
			Required: true,
			BodyPath: "protocol",
		},
		&requestflag.Flag[string]{
			Name:     "source-address",
			Usage:    "Source address of the Firewall Rule. Address of 0.0.0.0 requires a CIDR mask of 0. Must be in network-aligned/canonical form.",
			Required: true,
			BodyPath: "source_address",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Firewall Rule.",
			BodyPath: "tags",
		},
	},
	Action:          handleNetworkingFirewallRulesCreate,
	HideHelpCommand: true,
}

var networkingFirewallRulesUpdate = cli.Command{
	Name:    "update",
	Usage:   "Update a firewall rule",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "vpc-id",
			Required:  true,
			PathParam: "vpc_id",
		},
		&requestflag.Flag[string]{
			Name:      "firewall-rule-id",
			Required:  true,
			PathParam: "firewall_rule_id",
		},
		&requestflag.Flag[string]{
			Name:     "destination-address",
			Usage:    "Destination address of the Firewall Rule. Either VPC CIDR or VM in VPC. Must be in network-aligned/canonical form.",
			BodyPath: "destination_address",
		},
		&requestflag.Flag[[]string]{
			Name:     "destination-port",
			Usage:    "Destination ports of the Firewall Rule.",
			BodyPath: "destination_ports",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the Firewall Rule.",
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "protocol",
			Usage:    "Protocol of the Firewall Rule.",
			BodyPath: "protocol",
		},
		&requestflag.Flag[string]{
			Name:     "source-address",
			Usage:    "Source address of the Firewall Rule. Address of 0.0.0.0 requires a CIDR mask of 0. Must be in network-aligned/canonical form.",
			BodyPath: "source_address",
		},
		&requestflag.Flag[[]string]{
			Name:     "tag",
			Usage:    "Tags to attach to the Firewall Rule.",
			BodyPath: "tags",
		},
	},
	Action:          handleNetworkingFirewallRulesUpdate,
	HideHelpCommand: true,
}

var networkingFirewallRulesList = cli.Command{
	Name:    "list",
	Usage:   "List all firewall rules",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "vpc-id",
			Required:  true,
			PathParam: "vpc_id",
		},
		&requestflag.Flag[string]{
			Name:      "cursor",
			Usage:     "Pagination cursor returned by a previous request. Only valid for the same VPC, filters and sort order.",
			QueryPath: "cursor",
		},
		&requestflag.Flag[int64]{
			Name:      "limit",
			Usage:     "Maximum number of items to return",
			Default:   10,
			QueryPath: "limit",
		},
		&requestflag.Flag[string]{
			Name:      "name",
			Usage:     "Filter by a case-insensitive substring of the Firewall Rule name",
			QueryPath: "name",
		},
		&requestflag.Flag[string]{
			Name:      "protocol",
			Usage:     "Filter by protocol",
			QueryPath: "protocol",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, updated_at, name, status, protocol",
			Default:   "created_at:desc",
			QueryPath: "sort",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by Firewall Rule status",
			QueryPath: "status",
		},
		&requestflag.Flag[[]string]{
			Name:      "tag",
			Usage:     "Filter by tags. Repeat the parameter to require several tags; a Firewall Rule must carry all of them.",
			QueryPath: "tags",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleNetworkingFirewallRulesList,
	HideHelpCommand: true,
}

var networkingFirewallRulesDelete = cli.Command{
	Name:    "delete",
	Usage:   "Delete a firewall rule",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "vpc-id",
			Required:  true,
			PathParam: "vpc_id",
		},
		&requestflag.Flag[string]{
			Name:      "firewall-rule-id",
			Required:  true,
			PathParam: "firewall_rule_id",
		},
	},
	Action:          handleNetworkingFirewallRulesDelete,
	HideHelpCommand: true,
}

var networkingFirewallRulesGet = cli.Command{
	Name:    "get",
	Usage:   "Get details about a firewall rule",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "vpc-id",
			Required:  true,
			PathParam: "vpc_id",
		},
		&requestflag.Flag[string]{
			Name:      "firewall-rule-id",
			Required:  true,
			PathParam: "firewall_rule_id",
		},
	},
	Action:          handleNetworkingFirewallRulesGet,
	HideHelpCommand: true,
}

func handleNetworkingFirewallRulesCreate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vpc-id") && len(unusedArgs) > 0 {
		cmd.Set("vpc-id", unusedArgs[0])
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

	params := networking.FirewallRuleNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Networking.FirewallRules.New(
		ctx,
		cmd.Value("vpc-id").(string),
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
		Title:          "networking:firewall-rules create",
		Transform:      transform,
	})
}

func handleNetworkingFirewallRulesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vpc-id") && len(unusedArgs) > 0 {
		cmd.Set("vpc-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("firewall-rule-id") && len(unusedArgs) > 0 {
		cmd.Set("firewall-rule-id", unusedArgs[0])
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

	params := networking.FirewallRuleUpdateParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Networking.FirewallRules.Update(
		ctx,
		cmd.Value("vpc-id").(string),
		cmd.Value("firewall-rule-id").(string),
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
		Title:          "networking:firewall-rules update",
		Transform:      transform,
	})
}

func handleNetworkingFirewallRulesList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vpc-id") && len(unusedArgs) > 0 {
		cmd.Set("vpc-id", unusedArgs[0])
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

	params := networking.FirewallRuleListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Networking.FirewallRules.List(
			ctx,
			cmd.Value("vpc-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "networking:firewall-rules list",
			Transform:      transform,
		})
	} else {
		iter := client.Networking.FirewallRules.ListAutoPaging(
			ctx,
			cmd.Value("vpc-id").(string),
			params,
			options...,
		)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "networking:firewall-rules list",
			Transform:      transform,
		})
	}
}

func handleNetworkingFirewallRulesDelete(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vpc-id") && len(unusedArgs) > 0 {
		cmd.Set("vpc-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("firewall-rule-id") && len(unusedArgs) > 0 {
		cmd.Set("firewall-rule-id", unusedArgs[0])
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
	_, err = client.Networking.FirewallRules.Delete(
		ctx,
		cmd.Value("vpc-id").(string),
		cmd.Value("firewall-rule-id").(string),
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
		Title:          "networking:firewall-rules delete",
		Transform:      transform,
	})
}

func handleNetworkingFirewallRulesGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vpc-id") && len(unusedArgs) > 0 {
		cmd.Set("vpc-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("firewall-rule-id") && len(unusedArgs) > 0 {
		cmd.Set("firewall-rule-id", unusedArgs[0])
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
	_, err = client.Networking.FirewallRules.Get(
		ctx,
		cmd.Value("vpc-id").(string),
		cmd.Value("firewall-rule-id").(string),
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
		Title:          "networking:firewall-rules get",
		Transform:      transform,
	})
}
