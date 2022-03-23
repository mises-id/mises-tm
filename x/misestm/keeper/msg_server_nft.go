package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	nft "github.com/cosmos/cosmos-sdk/x/nft"
	"github.com/mises-id/mises-tm/x/misestm/types"
)

func (k msgServer) NewDenom(goCtx context.Context, msg *types.MsgNewDenom) (*types.MsgNewDenomResponse, error) {

	return &types.MsgNewDenomResponse{}, nil
}

func (k msgServer) NewNFTClass(goCtx context.Context, msg *types.MsgNewNFTClass) (*types.MsgNewNFTClassResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var nftClass = nft.Class{
		Id:          msg.Id,
		Name:        msg.Name,
		Symbol:      msg.Symbol,
		Uri:         msg.Uri,
		UriHash:     "",
		Data:        msg.Data,
		Description: "",
	}
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	nk := k.nk

	if err := nk.SaveClass(ctx, nftClass); err != nil {
		return nil, err
	}

	nk.SetClassOwner(ctx, msg.Id, sender)
	return &types.MsgNewNFTClassResponse{}, nil
}

func (k msgServer) UpdateNFTClass(goCtx context.Context, msg *types.MsgUpdateNFTClass) (*types.MsgUpdateNFTClassResponse, error) {

	return &types.MsgUpdateNFTClassResponse{}, nil
}

func (k msgServer) MintNFT(goCtx context.Context, msg *types.MsgMintNFT) (*types.MsgMintNFTResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var nft = nft.NFT{
		Id:      msg.Id,
		ClassId: msg.ClassId,
		Uri:     msg.Uri,
		UriHash: "",
		Data:    msg.Data,
	}
	nk := k.nk
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	owner := nk.GetClassOwner(ctx, msg.ClassId)
	if !owner.Equals(sender) {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnauthorized, "%s is not the owner of nft class %s", sender, msg.ClassId)
	}

	receiver, err := sdk.AccAddressFromBech32(msg.Recipient)
	if err != nil {
		return nil, err
	}

	if err := nk.Mint(ctx, nft, receiver); err != nil {
		return nil, err
	}

	return &types.MsgMintNFTResponse{}, nil
}

func (k msgServer) UpdateNFT(goCtx context.Context, msg *types.MsgUpdateNFT) (*types.MsgUpdateNFTResponse, error) {

	return &types.MsgUpdateNFTResponse{}, nil
}

func (k msgServer) BurnNFT(goCtx context.Context, msg *types.MsgBurnNFT) (*types.MsgBurnNFTResponse, error) {

	return &types.MsgBurnNFTResponse{}, nil
}
