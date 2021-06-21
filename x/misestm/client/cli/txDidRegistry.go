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

func CmdCreateDidRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-DidRegistry [did] [pkeyDid] [pkeyType] [pkeyMultibase] [version]",
		Short: "Create a new DidRegistry",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsDid, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}
			argsPkeyDid, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}
			argsPkeyType, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}
			argsPkeyMultibase, err := cast.ToStringE(args[3])
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

			msg := types.NewMsgCreateDidRegistry(clientCtx.GetFromAddress().String(), argsDid, argsPkeyDid, argsPkeyType, argsPkeyMultibase, argsVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateDidRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-DidRegistry [id] [did] [pkeyDid] [pkeyType] [pkeyMultibase] [version]",
		Short: "Update a DidRegistry",
		Args:  cobra.ExactArgs(6),
		RunE: func(cmd *cobra.Command, args []string) error {
			id, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			argsDid, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsPkeyDid, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsPkeyType, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			argsPkeyMultibase, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			argsVersion, err := cast.ToUint64E(args[5])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDidRegistry(clientCtx.GetFromAddress().String(), id, argsDid, argsPkeyDid, argsPkeyType, argsPkeyMultibase, argsVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdDeleteDidRegistry() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "delete-DidRegistry [id]",
		Short: "Delete a DidRegistry by id",
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

			msg := types.NewMsgDeleteDidRegistry(clientCtx.GetFromAddress().String(), id)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
