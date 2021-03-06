package cli

import (
	"fmt"
	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	// "github.com/cosmos/cosmos-sdk/client/flags"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/mises-id/mises-tm/x/misestm/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string) *cobra.Command {
	// Group misestm queries under a subcommand
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	// this line is used by starport scaffolding # 1

	cmd.AddCommand(CmdListUserInfo())
	cmd.AddCommand(CmdShowUserInfo())

	cmd.AddCommand(CmdListUserRelation())
	cmd.AddCommand(CmdShowUserRelation())

	cmd.AddCommand(CmdListAppInfo())
	cmd.AddCommand(CmdShowAppInfo())

	cmd.AddCommand(CmdListDidRegistry())
	cmd.AddCommand(CmdShowDidRegistry())

	return cmd
}
