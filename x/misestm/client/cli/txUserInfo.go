package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func CmdUpdateUserInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-UserInfo [uid] [encData] [iv] [version]",
		Short: "Update a UserInfo",
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

			msg := types.NewMsgUpdateUserInfo(clientCtx.GetFromAddress().String(), argsUid, argsEncData, argsIv, argsVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
