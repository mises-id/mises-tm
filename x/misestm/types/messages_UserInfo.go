package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateUserInfo{}

func NewMsgUpdateUserInfo(creator string, uid string, encData string, iv string, version uint64) *MsgUpdateUserInfo {
	return &MsgUpdateUserInfo{
		Creator: creator,
		Uid:     uid,
		PriInfo: &PrivateUserInfo{
			EncData: encData,
			Iv:      iv,
		},

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
