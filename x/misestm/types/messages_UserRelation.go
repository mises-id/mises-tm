package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateUserRelation{}

func NewMsgCreateUserRelation(creator string, uidFrom string, uidTo string, relType uint64, version uint64) *MsgCreateUserRelation {
	return &MsgCreateUserRelation{
		Creator: creator,
		UidFrom: uidFrom,
		UidTo:   uidTo,
		RelType: relType,
		Version: version,
	}
}

func (msg *MsgCreateUserRelation) Route() string {
	return RouterKey
}

func (msg *MsgCreateUserRelation) Type() string {
	return "CreateUserRelation"
}

func (msg *MsgCreateUserRelation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateUserRelation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateUserRelation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUserRelation{}

func NewMsgUpdateUserRelation(creator string, id uint64, uidFrom string, uidTo string, relType uint64, version uint64) *MsgUpdateUserRelation {
	return &MsgUpdateUserRelation{
		Id:      id,
		Creator: creator,
		UidFrom: uidFrom,
		UidTo:   uidTo,
		RelType: relType,
		Version: version,
	}
}

func (msg *MsgUpdateUserRelation) Route() string {
	return RouterKey
}

func (msg *MsgUpdateUserRelation) Type() string {
	return "UpdateUserRelation"
}

func (msg *MsgUpdateUserRelation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateUserRelation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUserRelation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUserRelation{}

func NewMsgDeleteUserRelation(creator string, id uint64) *MsgDeleteUserRelation {
	return &MsgDeleteUserRelation{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteUserRelation) Route() string {
	return RouterKey
}

func (msg *MsgDeleteUserRelation) Type() string {
	return "DeleteUserRelation"
}

func (msg *MsgDeleteUserRelation) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteUserRelation) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteUserRelation) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
