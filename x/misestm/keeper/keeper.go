package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/types"
	dbm "github.com/tendermint/tm-db"
	// this line is used by starport scaffolding # ibc/keeper/import
)

type (
	Keeper struct {
		cdc      codec.Codec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		ak       types.AccountKeeper
		fk       types.FeeGrantKeeper
		nk       types.NFTKeeper
		db       dbm.RawDB
		// this line is used by starport scaffolding # ibc/keeper/attribute
	}
)

func NewKeeper(
	cdc codec.Codec,
	storeKey,
	memKey sdk.StoreKey,
	ak types.AccountKeeper,
	fk types.FeeGrantKeeper,
	nk types.NFTKeeper,
	db dbm.RawDB,
	// this line is used by starport scaffolding # ibc/keeper/parameter
) *Keeper {

	k := &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
		ak:       ak,
		fk:       fk,
		nk:       nk,
		db:       db,
		// this line is used by starport scaffolding # ibc/keeper/return
	}

	return k
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
