package types

import (

	// nolint: staticcheck

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// baseAppSimulateFn is the signature of the Baseapp#Simulate function.
type baseAppSimulateFn func(txBytes []byte) (sdk.GasInfo, *sdk.Result, error)

var Simulater baseAppSimulateFn

func RegisterBaseAppSimulateFn(
	simulateFn baseAppSimulateFn,
) {
	Simulater = simulateFn
}
