package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAppInfo{}

func NewMsgCreateAppInfo(creator string, appid string, name string, domains []string, developer string, homeUrl string, iconUrl string, version uint64) *MsgCreateAppInfo {
	return &MsgCreateAppInfo{
		Creator:   creator,
		Appid:     appid,
		Name:      name,
		Domains:   domains,
		Developer: developer,
		HomeUrl:   homeUrl,
		IconUrl:   iconUrl,
		Version:   version,
	}
}

func (msg *MsgCreateAppInfo) Route() string {
	return RouterKey
}

func (msg *MsgCreateAppInfo) Type() string {
	return "CreateAppInfo"
}

func (msg *MsgCreateAppInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateAppInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateAppInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateAppInfo{}

func NewMsgUpdateAppInfo(creator string, id uint64, appid string, name string, domains []string, developer string, homeUrl string, iconUrl string, version uint64) *MsgUpdateAppInfo {
	return &MsgUpdateAppInfo{
		Id:        id,
		Creator:   creator,
		Appid:     appid,
		Name:      name,
		Domains:   domains,
		Developer: developer,
		HomeUrl:   homeUrl,
		IconUrl:   iconUrl,
		Version:   version,
	}
}

func (msg *MsgUpdateAppInfo) Route() string {
	return RouterKey
}

func (msg *MsgUpdateAppInfo) Type() string {
	return "UpdateAppInfo"
}

func (msg *MsgUpdateAppInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateAppInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateAppInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAppInfo{}

func NewMsgDeleteAppInfo(creator string, id uint64) *MsgDeleteAppInfo {
	return &MsgDeleteAppInfo{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteAppInfo) Route() string {
	return RouterKey
}

func (msg *MsgDeleteAppInfo) Type() string {
	return "DeleteAppInfo"
}

func (msg *MsgDeleteAppInfo) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteAppInfo) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteAppInfo) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
