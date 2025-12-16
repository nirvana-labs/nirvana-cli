// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command *cli.Command
)

func init() {
	Command = &cli.Command{
		Name:    "nirvana",
		Usage:   "CLI for the Nirvana Labs API",
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "user",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&userGet,
				},
			},
			{
				Name:     "api-keys",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&apiKeysCreate,
					&apiKeysUpdate,
					&apiKeysList,
					&apiKeysGet,
				},
			},
			{
				Name:     "operations",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&operationsList,
					&operationsGet,
				},
			},
			{
				Name:     "projects",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&projectsCreate,
					&projectsUpdate,
					&projectsList,
					&projectsGet,
				},
			},
			{
				Name:     "regions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&regionsList,
				},
			},
			{
				Name:     "compute:vms",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&computeVMsCreate,
					&computeVMsUpdate,
					&computeVMsList,
					&computeVMsDelete,
					&computeVMsGet,
					&computeVMsRestart,
				},
			},
			{
				Name:     "compute:vms:volumes",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&computeVMsVolumesList,
				},
			},
			{
				Name:     "compute:vms:os-images",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&computeVMsOSImagesList,
				},
			},
			{
				Name:     "compute:volumes",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&computeVolumesCreate,
					&computeVolumesUpdate,
					&computeVolumesList,
					&computeVolumesDelete,
					&computeVolumesGet,
				},
			},
			{
				Name:     "networking:vpcs",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&networkingVPCsCreate,
					&networkingVPCsUpdate,
					&networkingVPCsList,
					&networkingVPCsDelete,
					&networkingVPCsGet,
				},
			},
			{
				Name:     "networking:firewall-rules",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&networkingFirewallRulesCreate,
					&networkingFirewallRulesUpdate,
					&networkingFirewallRulesList,
					&networkingFirewallRulesDelete,
					&networkingFirewallRulesGet,
				},
			},
			{
				Name:     "networking:connect:connections",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&networkingConnectConnectionsCreate,
					&networkingConnectConnectionsUpdate,
					&networkingConnectConnectionsList,
					&networkingConnectConnectionsDelete,
					&networkingConnectConnectionsGet,
				},
			},
			{
				Name:     "networking:connect:routes",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&networkingConnectRoutesList,
				},
			},
			{
				Name:     "rpc-nodes:flex",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&rpcNodesFlexCreate,
					&rpcNodesFlexUpdate,
					&rpcNodesFlexList,
					&rpcNodesFlexGet,
				},
			},
			{
				Name:     "rpc-nodes:flex:blockchains",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&rpcNodesFlexBlockchainsList,
				},
			},
			{
				Name:     "rpc-nodes:dedicated",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&rpcNodesDedicatedList,
					&rpcNodesDedicatedGet,
				},
			},
			{
				Name:     "rpc-nodes:dedicated:blockchains",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&rpcNodesDedicatedBlockchainsList,
				},
			},
			{
				Name:     "vektor:registry:assets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorRegistryAssetsList,
				},
			},
			{
				Name:     "vektor:registry:blockchains",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorRegistryBlockchainsList,
				},
			},
			{
				Name:     "vektor:registry:venues",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorRegistryVenuesList,
				},
			},
			{
				Name:     "vektor:registry:errors",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorRegistryErrorsList,
				},
			},
			{
				Name:     "vektor:registry:lend-markets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorRegistryLendMarketsList,
				},
			},
			{
				Name:     "vektor:registry:borrow-markets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorRegistryBorrowMarketsList,
				},
			},
			{
				Name:     "vektor:registry:lp-pools",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorRegistryLPPoolsList,
				},
			},
			{
				Name:     "vektor:balances",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorBalancesList,
					&vektorBalancesListHistorical,
				},
			},
			{
				Name:     "vektor:prices",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorPricesList,
					&vektorPricesListHistorical,
				},
			},
			{
				Name:     "vektor:lend:markets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLendMarketsList,
					&vektorLendMarketsListHistorical,
				},
			},
			{
				Name:     "vektor:lend:positions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLendPositionsList,
					&vektorLendPositionsListHistorical,
				},
			},
			{
				Name:     "vektor:lend:lend",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLendLendCreate,
				},
			},
			{
				Name:     "vektor:lend:withdraw",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLendWithdrawCreate,
				},
			},
			{
				Name:     "vektor:lend:set-collateral",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLendSetCollateralCreate,
				},
			},
			{
				Name:     "vektor:borrow:markets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorBorrowMarketsList,
					&vektorBorrowMarketsListHistorical,
				},
			},
			{
				Name:     "vektor:borrow:positions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorBorrowPositionsList,
					&vektorBorrowPositionsListHistorical,
				},
			},
			{
				Name:     "vektor:borrow:accounts",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorBorrowAccountsList,
					&vektorBorrowAccountsListHistorical,
				},
			},
			{
				Name:     "vektor:borrow:borrow",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorBorrowBorrowCreate,
				},
			},
			{
				Name:     "vektor:borrow:repay",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorBorrowRepayCreate,
				},
			},
			{
				Name:     "vektor:lp:pools",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLPPoolsList,
					&vektorLPPoolsListHistorical,
				},
			},
			{
				Name:     "vektor:lp:positions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLPPositionsList,
					&vektorLPPositionsListHistorical,
				},
			},
			{
				Name:     "vektor:lp:deposit-quote",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLPDepositQuoteCreate,
				},
			},
			{
				Name:     "vektor:lp:withdraw-quote",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLPWithdrawQuoteCreate,
				},
			},
			{
				Name:     "vektor:buy:quotes",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorBuyQuotesList,
				},
			},
			{
				Name:     "vektor:buy:buy",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorBuyBuyCreate,
				},
			},
			{
				Name:     "vektor:sell:quotes",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorSellQuotesList,
				},
			},
			{
				Name:     "vektor:sell:sell",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorSellSellCreate,
				},
			},
			{
				Name:     "vektor:move",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorMoveCreate,
				},
			},
			{
				Name:     "vektor:wrap:wrap",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorWrapWrapCreate,
				},
			},
			{
				Name:     "vektor:wrap:unwrap",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorWrapUnwrapCreate,
				},
			},
			{
				Name:     "vektor:bridge:quotes",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorBridgeQuotesList,
				},
			},
			{
				Name:     "vektor:lock:markets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLockMarketsList,
				},
			},
			{
				Name:     "vektor:lock:positions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorLockPositionsList,
				},
			},
			{
				Name:     "vektor:vote:markets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorVoteMarketsList,
				},
			},
			{
				Name:     "vektor:vote:rewards",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorVoteRewardsList,
				},
			},
			{
				Name:     "vektor:incentivize",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorIncentivizeList,
				},
			},
			{
				Name:     "vektor:executions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorExecutionsList,
					&vektorExecutionsGet,
				},
			},
			{
				Name:     "vektor:executions:steps",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&vektorExecutionsStepsGet,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "nirvana @manpages [-o nirvana.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
		},
		EnableShellCompletion:      true,
		ShellCompletionCommandName: "@completion",
		HideHelpCommand:            true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "nirvana.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "nirvana.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
