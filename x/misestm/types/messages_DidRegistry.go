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
