package misestm

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/mises-id/mises-tm/x/misestm/keeper"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	for _, elem := range genState.MisesAccountList {
		k.SetMisesAccount(ctx, *elem)
	}

	// this line is used by starport scaffolding # genesis/module/init
	// Set all the UserInfo
	for _, elem := range genState.UserInfoList {
		k.SetUserInfo(ctx, *elem)
	}

	// Set UserInfo count
	k.SetUserInfoCount(ctx, genState.UserInfoCount)

	// Set all the UserRelation
	for _, elem := range genState.UserRelationList {
		k.SetUserRelation(ctx, *elem)
	}

	// Set UserRelation count
	k.SetUserRelationCount(ctx, genState.UserRelationCount)

	// Set all the AppInfo
	for _, elem := range genState.AppInfoList {
		k.SetAppInfo(ctx, *elem)
	}

	// Set AppInfo count
	k.SetAppInfoCount(ctx, genState.AppInfoCount)

	// Set all the DidRegistry
	for _, elem := range genState.DidRegistryList {
		k.SetDidRegistry(ctx, *elem)
	}

	// Set DidRegistry count
	k.SetDidRegistryCount(ctx, genState.DidRegistryCount)

	// this line is used by starport scaffolding # ibc/genesis/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	MisesAccountList := k.GetAllMisesAccount(ctx)
	for _, elem := range MisesAccountList {
		elem := elem
		genesis.MisesAccountList = append(genesis.MisesAccountList, &elem)
	}

	// this line is used by starport scaffolding # genesis/module/export
	// Get all UserInfo
	UserInfoList := k.GetAllUserInfo(ctx)
	for _, elem := range UserInfoList {
		elem := elem
		genesis.UserInfoList = append(genesis.UserInfoList, &elem)
	}

	// Set the current count
	genesis.UserInfoCount = k.GetUserInfoCount(ctx)

	// Get all UserRelation
	UserRelationList := k.GetAllUserRelation(ctx)
	for _, elem := range UserRelationList {
		elem := elem
		genesis.UserRelationList = append(genesis.UserRelationList, &elem)
	}

	// Set the current count
	genesis.UserRelationCount = k.GetUserRelationCount(ctx)

	// Get all AppInfo
	AppInfoList := k.GetAllAppInfo(ctx)
	for _, elem := range AppInfoList {
		elem := elem
		genesis.AppInfoList = append(genesis.AppInfoList, &elem)
	}

	// Set the current count
	genesis.AppInfoCount = k.GetAppInfoCount(ctx)

	// Get all DidRegistry
	DidRegistryList := k.GetAllDidRegistry(ctx)
	for _, elem := range DidRegistryList {
		elem := elem
		genesis.DidRegistryList = append(genesis.DidRegistryList, &elem)
	}

	// Set the current count
	genesis.DidRegistryCount = k.GetDidRegistryCount(ctx)

	// this line is used by starport scaffolding # ibc/genesis/export

	return genesis
}
