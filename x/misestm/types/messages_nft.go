package types

import (
	types1 "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

var _ sdk.Msg = &MsgNewDenom{}

func NewMsgNewDenom(id string, amount sdk.Int, denomMeta *banktypes.Metadata, sender string, recipient string) *MsgNewDenom {
	return &MsgNewDenom{
		Id:        id,
		Amount:    amount,
		DenomMeta: denomMeta,
		Sender:    sender,
		Recipient: recipient,
	}
}

func (msg *MsgNewDenom) Route() string {
	return RouterKey
}

func (msg *MsgNewDenom) Type() string {
	return "NewDenom"
}

func (msg *MsgNewDenom) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgNewDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewDenom) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgNewNFTClass{}

func NewMsgNewNFTClass(id string, name string, uri string, schema string, symbol string, data *types1.Any, sender string) *MsgNewNFTClass {
	return &MsgNewNFTClass{
		Id:     id,
		Name:   name,
		Uri:    uri,
		Schema: schema,
		Symbol: symbol,
		Data:   data,
		Sender: sender,
	}
}

func (msg *MsgNewNFTClass) Route() string {
	return RouterKey
}

func (msg *MsgNewNFTClass) Type() string {
	return "NewNFTClass"
}

func (msg *MsgNewNFTClass) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgNewNFTClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgNewNFTClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateNFTClass{}

func NewMsgUpdateNFTClass(id string, classId string, name string, uri string, data *types1.Any, sender string) *MsgUpdateNFTClass {
	return &MsgUpdateNFTClass{
		Id:      id,
		ClassId: classId,
		Name:    name,
		Uri:     uri,
		Data:    data,
		Sender:  sender,
	}
}

func (msg *MsgUpdateNFTClass) Route() string {
	return RouterKey
}

func (msg *MsgUpdateNFTClass) Type() string {
	return "UpdateNFTClass"
}

func (msg *MsgUpdateNFTClass) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgUpdateNFTClass) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateNFTClass) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgMintNFT{}

func NewMsgMintNFT(id string, classId string, name string, uri string, data *types1.Any, sender string, recipient string) *MsgMintNFT {
	return &MsgMintNFT{
		Id:        id,
		ClassId:   classId,
		Name:      name,
		Uri:       uri,
		Data:      data,
		Sender:    sender,
		Recipient: recipient,
	}
}

func (msg *MsgMintNFT) Route() string {
	return RouterKey
}

func (msg *MsgMintNFT) Type() string {
	return "MintNFT"
}

func (msg *MsgMintNFT) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgMintNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgMintNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateNFT{}

func NewMsgUpdateNFT(id string, classId string, uri string, data *types1.Any, sender string) *MsgUpdateNFT {
	return &MsgUpdateNFT{
		Id:      id,
		ClassId: classId,
		Uri:     uri,
		Data:    data,
		Sender:  sender,
	}
}

func (msg *MsgUpdateNFT) Route() string {
	return RouterKey
}

func (msg *MsgUpdateNFT) Type() string {
	return "UpdateNFT"
}

func (msg *MsgUpdateNFT) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgUpdateNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgBurnNFT{}

func NewMsgBurnNFT(id string, classId string, sender string) *MsgBurnNFT {
	return &MsgBurnNFT{
		Id:      id,
		ClassId: classId,
		Sender:  sender,
	}
}

func (msg *MsgBurnNFT) Route() string {
	return RouterKey
}

func (msg *MsgBurnNFT) Type() string {
	return "BurnNFT"
}

func (msg *MsgBurnNFT) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg *MsgBurnNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBurnNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
