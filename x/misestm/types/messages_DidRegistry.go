package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateDidRegistry{}

func NewMsgCreateDidRegistry(creator string, did string, pkeyDid string, pkeyType string, pkeyMultibase string, version uint64) *MsgCreateDidRegistry {
	return &MsgCreateDidRegistry{
		Creator:       creator,
		Did:           did,
		PkeyDid:       pkeyDid,
		PkeyType:      pkeyType,
		PkeyMultibase: pkeyMultibase,
		Version:       version,
	}
}

func (msg *MsgCreateDidRegistry) Route() string {
	return RouterKey
}

func (msg *MsgCreateDidRegistry) Type() string {
	return "CreateDidRegistry"
}

func (msg *MsgCreateDidRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateDidRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateDidRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateDidRegistry{}

func NewMsgUpdateDidRegistry(creator string, id uint64, did string, pkeyDid string, pkeyType string, pkeyMultibase string, version uint64) *MsgUpdateDidRegistry {
	return &MsgUpdateDidRegistry{
		Id:            id,
		Creator:       creator,
		Did:           did,
		PkeyDid:       pkeyDid,
		PkeyType:      pkeyType,
		PkeyMultibase: pkeyMultibase,
		Version:       version,
	}
}

func (msg *MsgUpdateDidRegistry) Route() string {
	return RouterKey
}

func (msg *MsgUpdateDidRegistry) Type() string {
	return "UpdateDidRegistry"
}

func (msg *MsgUpdateDidRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateDidRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateDidRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteDidRegistry{}

func NewMsgDeleteDidRegistry(creator string, id uint64) *MsgDeleteDidRegistry {
	return &MsgDeleteDidRegistry{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteDidRegistry) Route() string {
	return RouterKey
}

func (msg *MsgDeleteDidRegistry) Type() string {
	return "DeleteDidRegistry"
}

func (msg *MsgDeleteDidRegistry) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteDidRegistry) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteDidRegistry) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
