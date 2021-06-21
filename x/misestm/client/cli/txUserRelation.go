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

func CmdCreateUserRelation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-UserRelation [uidFrom] [uidTo] [relType] [version]",
		Short: "Create a new UserRelation",
		Args:  cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsUidFrom, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsUidTo, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsRelType, err := cast.ToUint64E(args[2])
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

			msg := types.NewMsgCreateUserRelation(clientCtx.GetFromAddress().String(), argsUidFrom, argsUidTo, argsRelType, argsVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateUserRelation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-UserRelation [id] [uidFrom] [uidTo] [relType] [version]",
		Short: "Update a UserRelation",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsUidFrom, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsUidTo, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsRelType, err := cast.ToUint64E(args[3])
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

			msg := types.NewMsgUpdateUserRelation(clientCtx.GetFromAddress().String(), id, argsUidFrom, argsUidTo, argsRelType, argsVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteUserRelation() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-UserRelation [id]",
		Short: "Delete a UserRelation by id",
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

			msg := types.NewMsgDeleteUserRelation(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
