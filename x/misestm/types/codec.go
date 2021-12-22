package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgNewDenom{}, "misestm/NewDenom", nil)
	cdc.RegisterConcrete(&MsgNewNFTClass{}, "misestm/NewNFTClass", nil)
	cdc.RegisterConcrete(&MsgUpdateNFTClass{}, "misestm/UpdateNFTClass", nil)
	cdc.RegisterConcrete(&MsgMintNFT{}, "misestm/MintNFT", nil)
	cdc.RegisterConcrete(&MsgUpdateNFT{}, "misestm/UpdateNFT", nil)
	cdc.RegisterConcrete(&MsgBurnNFT{}, "misestm/BurnNFT", nil)

	// this line is used by starport scaffolding # 2

	cdc.RegisterConcrete(&MsgUpdateUserInfo{}, "misestm/UpdateUserInfo", nil)
	cdc.RegisterConcrete(&MsgUpdateUserRelation{}, "misestm/UpdateUserRelation", nil)
	cdc.RegisterConcrete(&MsgUpdateAppInfo{}, "misestm/UpdateAppInfo", nil)
	cdc.RegisterConcrete(&MsgCreateDidRegistry{}, "misestm/CreateDidRegistry", nil)

}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgNewDenom{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgNewNFTClass{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateNFTClass{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgMintNFT{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateNFT{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBurnNFT{},
	)

	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateUserInfo{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgUpdateUserRelation{},
	)
	registry.RegisterImplementations((*sdk.Msg)(nil))
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateDidRegistry{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)
