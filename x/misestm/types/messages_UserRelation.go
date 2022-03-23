package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateUserRelation{}

func NewMsgUpdateUserRelation(creator string, uidFrom string, uidTo string, relType uint64, version uint64) *MsgUpdateUserRelation {
	return &MsgUpdateUserRelation{
		Creator:      creator,
		UidFrom:      uidFrom,
		UidTo:        uidTo,
		IsFollowing:  relType&RelTypeBitFollow != 0,
		IsBlocking:   relType&RelTypeBitBlock != 0,
		IsReferredBy: relType&RelTypeBitReferredBy != 0,
		Version:      version,
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
