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

	"github.com/nirvana-labs/nirvana-cli/internal/autocomplete"
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
		Suggest: true,
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
				Suggest:  true,
				Commands: []*cli.Command{
					&userGet,
				},
			},
			{
				Name:     "user:security",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&userSecurityUpdate,
					&userSecurityGet,
				},
			},
			{
				Name:     "api-keys",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&apiKeysCreate,
					&apiKeysUpdate,
					&apiKeysList,
					&apiKeysDelete,
					&apiKeysGet,
				},
			},
			{
				Name:     "operations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&operationsList,
					&operationsGet,
				},
			},
			{
				Name:     "organizations",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&organizationsList,
					&organizationsGet,
				},
			},
			{
				Name:     "projects",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&projectsCreate,
					&projectsUpdate,
					&projectsList,
					&projectsDelete,
					&projectsGet,
				},
			},
			{
				Name:     "regions",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&regionsList,
					&regionsGet,
				},
			},
			{
				Name:     "compute:vms",
				Category: "API RESOURCE",
				Suggest:  true,
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
				Suggest:  true,
				Commands: []*cli.Command{
					&computeVMsVolumesList,
				},
			},
			{
				Name:     "compute:vms:os-images",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&computeVMsOSImagesList,
				},
			},
			{
				Name:     "compute:volumes",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&computeVolumesCreate,
					&computeVolumesUpdate,
					&computeVolumesList,
					&computeVolumesDelete,
					&computeVolumesAttach,
					&computeVolumesDetach,
					&computeVolumesGet,
				},
			},
			{
				Name:     "networking:vpcs",
				Category: "API RESOURCE",
				Suggest:  true,
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
				Suggest:  true,
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
				Suggest:  true,
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
				Suggest:  true,
				Commands: []*cli.Command{
					&networkingConnectRoutesList,
				},
			},
			{
				Name:     "rpc-nodes:flex",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&rpcNodesFlexCreate,
					&rpcNodesFlexUpdate,
					&rpcNodesFlexList,
					&rpcNodesFlexDelete,
					&rpcNodesFlexGet,
				},
			},
			{
				Name:     "rpc-nodes:flex:blockchains",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&rpcNodesFlexBlockchainsList,
				},
			},
			{
				Name:     "rpc-nodes:dedicated",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&rpcNodesDedicatedList,
					&rpcNodesDedicatedGet,
				},
			},
			{
				Name:     "rpc-nodes:dedicated:blockchains",
				Category: "API RESOURCE",
				Suggest:  true,
				Commands: []*cli.Command{
					&rpcNodesDedicatedBlockchainsList,
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
			{
				Name:            "__complete",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.ExecuteShellCompletion,
			},
			{
				Name:            "@completion",
				Hidden:          true,
				HideHelpCommand: true,
				Action:          autocomplete.OutputCompletionScript,
			},
		},
		HideHelpCommand: true,
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
