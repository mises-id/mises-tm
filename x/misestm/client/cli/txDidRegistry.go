package cli

import (
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
