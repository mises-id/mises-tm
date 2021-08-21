package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateUserInfo{}

func NewMsgCreateUserInfo(creator string, uid string, encData string, iv string, version uint64) *MsgCreateUserInfo {
	return &MsgCreateUserInfo{
		Creator: creator,
		Uid:     uid,
		EncData: encData,
		Iv:      iv,
		Version: version,
	}
}

func (msg *MsgCreateUserInfo) Route() string {
	return RouterKey
}

func (msg *MsgCreateUserInfo) Type() string {
	return "CreateUserInfo"
}

func (msg *MsgCreateUserInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateUserInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateUserInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateUserInfo{}

func NewMsgUpdateUserInfo(creator string, id uint64, uid string, encData string, iv string, version uint64) *MsgUpdateUserInfo {
	return &MsgUpdateUserInfo{
		Id:      id,
		Creator: creator,
		Uid:     uid,
		EncData: encData,
		Iv:      iv,
		Version: version,
	}
}

func (msg *MsgUpdateUserInfo) Route() string {
	return RouterKey
}

func (msg *MsgUpdateUserInfo) Type() string {
	return "UpdateUserInfo"
}

func (msg *MsgUpdateUserInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateUserInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateUserInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteUserInfo{}

func NewMsgDeleteUserInfo(creator string, id uint64) *MsgDeleteUserInfo {
	return &MsgDeleteUserInfo{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteUserInfo) Route() string {
	return RouterKey
}

func (msg *MsgDeleteUserInfo) Type() string {
	return "DeleteUserInfo"
}

func (msg *MsgDeleteUserInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteUserInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteUserInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
