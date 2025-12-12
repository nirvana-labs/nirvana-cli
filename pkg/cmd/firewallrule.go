// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/networking"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/stainless-sdks/nirvana-cli/internal/apiquery"
	"github.com/stainless-sdks/nirvana-cli/internal/requestflag"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var networkingFirewallRulesCreate = cli.Command{
	Name:  "create",
	Usage: "Create a firewall rule",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vpc-id",
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
	Action:          handleNetworkingFirewallRulesCreate,
	HideHelpCommand: true,
}

var networkingFirewallRulesUpdate = cli.Command{
	Name:  "update",
	Usage: "Update a firewall rule",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vpc-id",
		},
		&requestflag.Flag[string]{
			Name: "firewall-rule-id",
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
	Name:  "list",
	Usage: "List all firewall rules",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vpc-id",
		},
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
	Action:          handleNetworkingFirewallRulesList,
	HideHelpCommand: true,
}

var networkingFirewallRulesDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete a firewall rule",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vpc-id",
		},
		&requestflag.Flag[string]{
			Name: "firewall-rule-id",
		},
	},
	Action:          handleNetworkingFirewallRulesDelete,
	HideHelpCommand: true,
}

var networkingFirewallRulesGet = cli.Command{
	Name:  "get",
	Usage: "Get details about a firewall rule",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "vpc-id",
		},
		&requestflag.Flag[string]{
			Name: "firewall-rule-id",
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
	params := networking.FirewallRuleNewParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:firewall-rules create", obj, format, transform)
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
	params := networking.FirewallRuleUpdateParams{}

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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:firewall-rules update", obj, format, transform)
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
	params := networking.FirewallRuleListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
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
		return ShowJSON(os.Stdout, "networking:firewall-rules list", obj, format, transform)
	} else {
		iter := client.Networking.FirewallRules.ListAutoPaging(
			ctx,
			cmd.Value("vpc-id").(string),
			params,
			options...,
		)
		return streamOutput("networking:firewall-rules list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "networking:firewall-rules list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
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
		ApplicationJSON,
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:firewall-rules delete", obj, format, transform)
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
		ApplicationJSON,
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
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "networking:firewall-rules get", obj, format, transform)
}
