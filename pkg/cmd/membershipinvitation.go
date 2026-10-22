// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/nirvana-labs/nirvana-cli/internal/apiquery"
	"github.com/nirvana-labs/nirvana-cli/internal/requestflag"
	"github.com/nirvana-labs/nirvana-go/v2"
	"github.com/nirvana-labs/nirvana-go/v2/option"
	"github.com/nirvana-labs/nirvana-go/v2/organizations"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var organizationsMembershipsInvitationsCreate = cli.Command{
	Name:    "create",
	Usage:   "Invite an email address to join an organization with a role. The membership is\ncreated when the invitee accepts. Owners may invite with any role; admins may\ninvite with admin, billing or member.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[string]{
			Name:     "email",
			Usage:    "Email address to invite. Inviting an address that already belongs to a member of the organization fails.",
			Required: true,
			BodyPath: "email",
		},
		&requestflag.Flag[string]{
			Name:     "role",
			Usage:    "Role of the user in the organization.",
			Required: true,
			BodyPath: "role",
		},
	},
	Action:          handleOrganizationsMembershipsInvitationsCreate,
	HideHelpCommand: true,
}

var organizationsMembershipsInvitationsList = cli.Command{
	Name:    "list",
	Usage:   "List the invitations to an organization that are still pending: not yet\naccepted, revoked or expired. Only owners and admins may list invitations.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
	},
	Action:          handleOrganizationsMembershipsInvitationsList,
	HideHelpCommand: true,
}

var organizationsMembershipsInvitationsDelete = cli.Command{
	Name:    "delete",
	Usage:   "Revoke a pending invitation so it can no longer be accepted. Owners may revoke\nany invitation; admins may revoke invitations with the admin, billing or member\nrole.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[string]{
			Name:      "invitation-id",
			Required:  true,
			PathParam: "invitation_id",
		},
	},
	Action:          handleOrganizationsMembershipsInvitationsDelete,
	HideHelpCommand: true,
}

var organizationsMembershipsInvitationsResend = cli.Command{
	Name:    "resend",
	Usage:   "Send a pending invitation's email again. Owners may resend any invitation;\nadmins may resend invitations with the admin, billing or member role.",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "organization-id",
			Required:  true,
			PathParam: "organization_id",
		},
		&requestflag.Flag[string]{
			Name:      "invitation-id",
			Required:  true,
			PathParam: "invitation_id",
		},
	},
	Action:          handleOrganizationsMembershipsInvitationsResend,
	HideHelpCommand: true,
}

func handleOrganizationsMembershipsInvitationsCreate(ctx context.Context, cmd *cli.Command) error {
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

	params := organizations.MembershipInvitationNewParams{}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Organizations.Memberships.Invitations.New(
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
		Title:          "organizations:memberships:invitations create",
		Transform:      transform,
	})
}

func handleOrganizationsMembershipsInvitationsList(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Organizations.Memberships.Invitations.List(ctx, cmd.Value("organization-id").(string), options...)
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
		Title:          "organizations:memberships:invitations list",
		Transform:      transform,
	})
}

func handleOrganizationsMembershipsInvitationsDelete(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("organization-id") && len(unusedArgs) > 0 {
		cmd.Set("organization-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("invitation-id") && len(unusedArgs) > 0 {
		cmd.Set("invitation-id", unusedArgs[0])
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

	return client.Organizations.Memberships.Invitations.Delete(
		ctx,
		cmd.Value("organization-id").(string),
		cmd.Value("invitation-id").(string),
		options...,
	)
}

func handleOrganizationsMembershipsInvitationsResend(ctx context.Context, cmd *cli.Command) error {
	client := nirvana.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("organization-id") && len(unusedArgs) > 0 {
		cmd.Set("organization-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if !cmd.IsSet("invitation-id") && len(unusedArgs) > 0 {
		cmd.Set("invitation-id", unusedArgs[0])
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
	_, err = client.Organizations.Memberships.Invitations.Resend(
		ctx,
		cmd.Value("organization-id").(string),
		cmd.Value("invitation-id").(string),
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
		Title:          "organizations:memberships:invitations resend",
		Transform:      transform,
	})
}
