package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgCreateUserInfo{}, "misestm/CreateUserInfo", nil)
	cdc.RegisterConcrete(&MsgUpdateUserInfo{}, "misestm/UpdateUserInfo", nil)
	cdc.RegisterConcrete(&MsgDeleteUserInfo{}, "misestm/DeleteUserInfo", nil)

	cdc.RegisterConcrete(&MsgCreateUserRelation{}, "misestm/CreateUserRelation", nil)
	cdc.RegisterConcrete(&MsgUpdateUserRelation{}, "misestm/UpdateUserRelation", nil)
	cdc.RegisterConcrete(&MsgDeleteUserRelation{}, "misestm/DeleteUserRelation", nil)

	cdc.RegisterConcrete(&MsgCreateAppInfo{}, "misestm/CreateAppInfo", nil)
	cdc.RegisterConcrete(&MsgUpdateAppInfo{}, "misestm/UpdateAppInfo", nil)
	cdc.RegisterConcrete(&MsgDeleteAppInfo{}, "misestm/DeleteAppInfo", nil)

	cdc.RegisterConcrete(&MsgCreateDidRegistry{}, "misestm/CreateDidRegistry", nil)
	cdc.RegisterConcrete(&MsgUpdateDidRegistry{}, "misestm/UpdateDidRegistry", nil)
	cdc.RegisterConcrete(&MsgDeleteDidRegistry{}, "misestm/DeleteDidRegistry", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUserInfo{},
		&MsgUpdateUserInfo{},
		&MsgDeleteUserInfo{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateUserRelation{},
		&MsgUpdateUserRelation{},
		&MsgDeleteUserRelation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateAppInfo{},
		&MsgUpdateAppInfo{},
		&MsgDeleteAppInfo{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDidRegistry{},
		&MsgUpdateDidRegistry{},
		&MsgDeleteDidRegistry{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
