package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func CmdUpdateAppInfo() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-AppInfo [appid] [name]  [domain] [developer] [homeUrl] [iconUrl] [version]",
		Short: "update AppInfo",
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

			msg := types.NewMsgUpdateAppInfo(clientCtx.GetFromAddress().String(), argsAppid, argsName, []string{argsDomain}, argsDeveloper, argsHomeUrl, argsIconUrl, argsVersion)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
