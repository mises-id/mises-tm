package simapp

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmstate "github.com/tendermint/tendermint/proto/tendermint/state"
	sm "github.com/tendermint/tendermint/state"

	"github.com/tendermint/tendermint/types"
)

func TestABCIResponsesResultsHash(t *testing.T) {
	//"178D2E5F3880C9216B940B6807ABB332B995FEB9617A8CA5AD1FCD9BF35834CC"
	//"0F293D574F3638F58585A8EE0363AE5ED630A340A169307E4942545FBAF28CDA" 62444
	data, err := base64.StdEncoding.DecodeString("CiwKKi9taXNlc2lkLm1pc2VzdG0ubWlzZXN0bS5Nc2dVcGRhdGVVc2VySW5mbw==")
	responses := &tmstate.ABCIResponses{
		BeginBlock: &abci.ResponseBeginBlock{},
		DeliverTxs: []*abci.ResponseDeliverTx{
			{Code: 0, Data: data, Log: "Huh?", GasWanted: 100000, GasUsed: 62444},
		},
		EndBlock: &abci.ResponseEndBlock{},
	}

	root := sm.ABCIResponsesResultsHash(responses)
	fmt.Println(hex.EncodeToString(root))

	// root should be Merkle tree root of DeliverTxs responses
	results := types.NewResults(responses.DeliverTxs)
	assert.Equal(t, root, results.Hash())

	// test we can prove first DeliverTx
	proof := results.ProveResult(0)
	bz, err := results[0].Marshal()
	require.NoError(t, err)
	assert.NoError(t, proof.Verify(root, bz))
}
