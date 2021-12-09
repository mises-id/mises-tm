package cmd

import (
	"context"
	"testing"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	genutiltest "github.com/cosmos/cosmos-sdk/x/genutil/client/testutil"
	"github.com/mises-id/mises-tm/app"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
)

func Test_LightCmd(t *testing.T) {
	home := app.DefaultNodeHome
	app.SetConfig()
	encodingConfig := app.MakeEncodingConfig()
	logger := log.NewNopLogger()
	cfg, err := genutiltest.CreateDefaultTendermintConfig(home)
	require.NoError(t, err)

	serverCtx := server.NewContext(viper.New(), cfg, logger)
	clientCtx := client.Context{}.
		WithCodec(encodingConfig.Codec).
		WithHomeDir(home).
		WithTxConfig(encodingConfig.TxConfig).
		WithAccountRetriever(types.AccountRetriever{})

	ctx := context.Background()
	ctx = context.WithValue(ctx, server.ServerContextKey, serverCtx)
	ctx = context.WithValue(ctx, client.ClientContextKey, &clientCtx)
	cmd := LightCmd()
	cmd.SetArgs([]string{
		"test",
		"--listening-address=tcp://0.0.0.0:26657",
		"--log-level=trace",
		"--primary-addr=http://e1.mises.site:26657",
		"--witness-addr=http://e2.mises.site:26657",
		"--trusted-height=582507",
		"--trusted-hash=3F541BDF3CF2CE414FB4A3FAF90931101C4ABD31093239AC7E7A787B3E387230",
	})

	err = cmd.ExecuteContext(ctx)
	require.NoError(t, err)

}
