package cli

import (
	"context"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/mises-id/mises-tm/x/misestm/types"
	"github.com/spf13/cobra"
)

func CmdListDidRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-DidRegistry",
		Short: "list all DidRegistry",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllDidRegistryRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.DidRegistryAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddPaginationFlagsToCmd(cmd, cmd.Use)
	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowDidRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-DidRegistry [id]",
		Short: "shows a DidRegistry",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			params := &types.QueryGetDidRegistryRequest{
				Id: id,
			}

			res, err := queryClient.DidRegistry(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
