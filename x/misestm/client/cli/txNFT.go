package cli

import (
	"github.com/spf13/cobra"

	"github.com/spf13/cast"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func CmdCreateNFTClass() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create-nft-class [id] [name] [uri] [schema] [symbol]",
		Short: "Create a new NFT Class",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsId, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			argsName, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsUri, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsSchema, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}

			argsSymbol, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgNewNFTClass(argsId, argsName, argsUri, argsSchema, argsSymbol, nil, clientCtx.GetFromAddress().String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mint-nft [id] [classid] [name] [uri] [recipient]",
		Short: "mint a new NFT",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) error {
			argsId, err := cast.ToStringE(args[0])
			if err != nil {
				return err
			}

			argsClassId, err := cast.ToStringE(args[1])
			if err != nil {
				return err
			}

			argsName, err := cast.ToStringE(args[2])
			if err != nil {
				return err
			}

			argsUri, err := cast.ToStringE(args[3])
			if err != nil {
				return err
			}
			argsRecipient, err := cast.ToStringE(args[4])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := types.NewMsgMintNFT(argsId, argsClassId, argsName, argsUri, nil, clientCtx.GetFromAddress().String(), argsRecipient)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
