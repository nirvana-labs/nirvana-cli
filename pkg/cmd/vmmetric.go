// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go"
	"github.com/nirvana-labs/nirvana-go/compute"
	"github.com/nirvana-labs/nirvana-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var computeVMsMetricsList = cli.Command{
	Name:    "list",
	Usage:   "Read a VM's resource metrics over an interval. Every series covers the same\nperiods, so they line up index for index, and a period the VM reported no\nobservation for carries a null value.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "vm-id",
			Required:  true,
			PathParam: "vm_id",
		},
		&requestflag.Flag[string]{
			Name:      "aggregation",
			Usage:     "How the samples inside one period are folded into a single value.",
			Default:   "mean",
			QueryPath: "aggregation",
		},
		&requestflag.Flag[any]{
			Name:      "end-time",
			Usage:     "End of the interval, exclusive, as an RFC 3339 timestamp. Defaults to now.",
			QueryPath: "end_time",
		},
		&requestflag.Flag[[]string]{
			Name:      "metric",
			Usage:     "Metric to return. Repeat the parameter for several; every metric is returned when it is left out.",
			QueryPath: "metric",
		},
		&requestflag.Flag[string]{
			Name:      "period",
			Usage:     "Width of one period, and so the spacing between points. An interval holding more than 1440 periods is rejected; the error names a period that fits.",
			Default:   "5m",
			QueryPath: "period",
		},
		&requestflag.Flag[any]{
			Name:      "start-time",
			Usage:     "Start of the interval, inclusive, as an RFC 3339 timestamp. Defaults to an hour before end_time. A start older than the 30 days of history kept is served from where that history begins.",
			QueryPath: "start_time",
		},
	},
	Action:          handleComputeVMsMetricsList,
	HideHelpCommand: true,
}

func handleComputeVMsMetricsList(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("vm-id") && len(unusedArgs) > 0 {
		cmd.Set("vm-id", unusedArgs[0])
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

	params := compute.VMMetricListParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Compute.VMs.Metrics.List(
		ctx,
		cmd.Value("vm-id").(string),
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
		Title:          "compute:vms:metrics list",
		Transform:      transform,
	})
}
