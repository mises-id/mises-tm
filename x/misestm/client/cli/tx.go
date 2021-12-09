package cli

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

var (
	DefaultRelativePacketTimeoutTimestamp = uint64((time.Duration(10) * time.Minute).Nanoseconds())
)

const (
	flagPacketTimeoutTimestamp = "packet-timeout-timestamp"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdCreateNFTClass())
	cmd.AddCommand(CmdMintNFT())

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdCreateUserInfo())
	cmd.AddCommand(CmdUpdateUserInfo())
	cmd.AddCommand(CmdDeleteUserInfo())

	cmd.AddCommand(CmdCreateUserRelation())
	cmd.AddCommand(CmdUpdateUserRelation())
	cmd.AddCommand(CmdDeleteUserRelation())

	cmd.AddCommand(CmdCreateAppInfo())
	cmd.AddCommand(CmdUpdateAppInfo())
	cmd.AddCommand(CmdDeleteAppInfo())

	cmd.AddCommand(CmdCreateDidRegistry())
	cmd.AddCommand(CmdUpdateDidRegistry())
	cmd.AddCommand(CmdDeleteDidRegistry())

	return cmd
}
