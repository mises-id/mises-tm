package cli

import (
	"strconv"

	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func CmdCreateAppInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-AppInfo [appid] [name]  [domain] [developer] [homeUrl] [iconUrl] [version]",
		Short: "Create a new AppInfo",
		Args:  cobra.ExactArgs(7),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsAppid, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsName, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsDomain, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsDeveloper, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsHomeUrl, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}
			argsIconUrl, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}
			argsVersion, err := cast.ToUint64E(args[6])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateAppInfo(clientCtx.GetFromAddress().String(), argsAppid, argsName, []string{argsDomain}, argsDeveloper, argsHomeUrl, argsIconUrl, argsVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateAppInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-AppInfo [id] [appid] [name]  [domain] [developer] [homeUrl] [iconUrl] [version]",
		Short: "Update a AppInfo",
		Args:  cobra.ExactArgs(8),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsAppid, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsName, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsDomain, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsDeveloper, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}
			argsHomeUrl, err := cast.ToStringE(args[5])
			if err != nil {
				return err
			}
			argsIconUrl, err := cast.ToStringE(args[6])
			if err != nil {
				return err
			}
			argsVersion, err := cast.ToUint64E(args[7])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateAppInfo(clientCtx.GetFromAddress().String(), id, argsAppid, argsName, []string{argsDomain}, argsDeveloper, argsHomeUrl, argsIconUrl, argsVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteAppInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-AppInfo [id]",
		Short: "Delete a AppInfo by id",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgDeleteAppInfo(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
