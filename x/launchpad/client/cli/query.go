package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/nebula-labs/nebula/x/launchpad/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group launchpad queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdQueryParams())
	cmd.AddCommand(GetCmdQueryProject())
	cmd.AddCommand(GetCmdQueryTotalProject())
	cmd.AddCommand(GetCmdQueryProjectBalances())

	return cmd
}

func CmdQueryParams() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "params",
		Short: "shows the parameters of the module",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			res, err := queryClient.Params(context.Background(), &types.QueryParamsRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func GetCmdQueryProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project [project_id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query detail of project",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			projectId, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}
			// Query store
			res, err := queryClient.Project(context.Background(), &types.ProjectRequest{ProjectId: projectId.Uint64()})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	return cmd
}

func GetCmdQueryTotalProject() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "total-project",
		Args:  cobra.NoArgs,
		Short: "Query all projects",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			// Query store
			res, err := queryClient.TotalProject(context.Background(), &types.TotalProjectRequest{Pagination: pageReq})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)
	flags.AddPaginationFlagsToCmd(cmd, "launchpad")
	return cmd
}

func GetCmdQueryProjectBalances() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "project-balances [project-id]",
		Args:  cobra.ExactArgs(1),
		Short: "Query project balances",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := types.NewQueryClient(clientCtx)

			projectId, err := sdk.ParseUint(args[0])
			if err != nil {
				return err
			}

			// Query store
			res, err := queryClient.ProjectBalances(context.Background(), &types.ProjectBalancesRequest{
				ProjectId: projectId.Uint64(),
			})
			if err != nil {
				return err
			}
			return clientCtx.PrintProto(res)
		},
	}
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
