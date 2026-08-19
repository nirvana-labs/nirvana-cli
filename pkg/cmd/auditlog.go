// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/audit_logs"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var auditLogsList = cli.Command{
	Name:    "list",
	Usage:   "List Audit Log entries for an organization",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "action",
			Usage:     "Filter by recorded action",
			QueryPath: "action",
		},
		&requestflag.Flag[string]{
			Name:      "actor-id",
			Usage:     "Filter by the acting user or API key",
			QueryPath: "actor_id",
		},
		&requestflag.Flag[string]{
			Name:      "actor-type",
			Usage:     "Filter by the kind of actor that acted",
			QueryPath: "actor_type",
		},
		&requestflag.Flag[string]{
			Name:      "client-ip",
			Usage:     "Filter by client IP address, matched exactly",
			QueryPath: "client_ip",
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
			Name:      "method",
			Usage:     "Filter by HTTP method",
			QueryPath: "method",
		},
		&requestflag.Flag[string]{
			Name:      "path",
			Usage:     "Filter by a case-insensitive substring of the request path",
			QueryPath: "path",
		},
		&requestflag.Flag[string]{
			Name:      "sort",
			Usage:     "Comma-separated sort terms in precedence order, each field:asc or field:desc. Fields: created_at, status_code",
			Default:   "created_at:desc",
			QueryPath: "sort",
		},
		&requestflag.Flag[int64]{
			Name:      "status-code-max",
			Usage:     "Only entries with a status code at or below this",
			QueryPath: "status_code_max",
		},
		&requestflag.Flag[int64]{
			Name:      "status-code-min",
			Usage:     "Only entries with a status code at or above this, e.g. 400 for failures only",
			QueryPath: "status_code_min",
		},
		&requestflag.Flag[string]{
			Name:      "target-id",
			Usage:     "Filter by the resource acted on",
			QueryPath: "target_id",
		},
		&requestflag.Flag[string]{
			Name:      "target-type",
			Usage:     "Filter by the kind of resource acted on",
			QueryPath: "target_type",
		},
		&requestflag.Flag[int64]{
			Name:  "max-items",
			Usage: "The maximum number of items to return (use -1 for unlimited).",
		},
	},
	Action:          handleAuditLogsList,
	HideHelpCommand: true,
}

var auditLogsGet = cli.Command{
	Name:    "get",
	Usage:   "Get an Audit Log entry",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "audit-log-id",
			Required:  true,
			PathParam: "audit_log_id",
		},
	},
	Action:          handleAuditLogsGet,
	HideHelpCommand: true,
}

func handleAuditLogsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

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

	params := audit_logs.AuditLogListParams{}

	format := cmd.Root().String("format")
	explicitFormat := cmd.Root().IsSet("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.AuditLogs.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(obj, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "audit-logs list",
			Transform:      transform,
		})
	} else {
		iter := client.AuditLogs.ListAutoPaging(ctx, params, options...)
		maxItems := int64(-1)
		if cmd.IsSet("max-items") {
			maxItems = cmd.Value("max-items").(int64)
		}
		return ShowJSONIterator(iter, maxItems, ShowJSONOpts{
			ExplicitFormat: explicitFormat,
			Format:         format,
			RawOutput:      cmd.Root().Bool("raw-output"),
			Title:          "audit-logs list",
			Transform:      transform,
		})
	}
}

func handleAuditLogsGet(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("audit-log-id") && len(unusedArgs) > 0 {
		cmd.Set("audit-log-id", unusedArgs[0])
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
	_, err = client.AuditLogs.Get(ctx, cmd.Value("audit-log-id").(string), options...)
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
		Title:          "audit-logs get",
		Transform:      transform,
	})
}
