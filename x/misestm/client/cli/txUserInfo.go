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

func CmdCreateUserInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-UserInfo [uid] [encData] [iv] [version]",
		Short: "Create a new UserInfo",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsUid, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsEncData, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsIv, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsVersion, err := cast.ToUint64E(args[3])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateUserInfo(clientCtx.GetFromAddress().String(), argsUid, argsEncData, argsIv, argsVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateUserInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-UserInfo [id] [uid] [encData] [iv] [version]",
		Short: "Update a UserInfo",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsUid, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsEncData, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsIv, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			argsVersion, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateUserInfo(clientCtx.GetFromAddress().String(), id, argsUid, argsEncData, argsIv, argsVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteUserInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-UserInfo [id]",
		Short: "Delete a UserInfo by id",
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

			msg := types.NewMsgDeleteUserInfo(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
