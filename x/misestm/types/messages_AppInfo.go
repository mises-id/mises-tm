package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateAppInfo{}

func NewMsgUpdateAppInfo(creator string, appid string, name string, domains []string, developer string, homeUrl string, iconUrl string, version uint64) *MsgUpdateAppInfo {
	return &MsgUpdateAppInfo{
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
