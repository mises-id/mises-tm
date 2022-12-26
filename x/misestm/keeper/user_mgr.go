package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userMgr struct {
	Keeper
}

func NewUserMgrImpl(keeper Keeper) types.UserMgr {
	return &userMgr{Keeper: keeper}
}

var _ types.UserMgr = &userMgr{}

func (k *userMgr) GetUserAccount(ctx sdk.Context, did string) (*types.MisesAccount, error) {
	ak := k.ak
	addr, didtype, err := types.AddrFromDid(did)
	if err != nil {
		return nil, err
	}

	if acc := ak.GetAccount(ctx, addr); acc == nil {
		return nil, nil
	}
	if !k.HasMisesAccount(ctx, did) {
		return nil, sdkerrors.ErrKeyNotFound
	}
	misesAcc := k.GetMisesAccount(ctx, did)

	if didtype != misesAcc.DidType || didtype != types.DIDTypeUser {
		return nil, sdkerrors.ErrLogic
	}
	return &misesAcc, nil
}

func (k *userMgr) GetUserRelation(ctx sdk.Context, didFrom string, didTo string) (ret *types.UserRelation, err error) {
	if didFrom == didTo {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "from user must diff from to user")
	}
	misesAccFrom, err := k.GetUserAccount(ctx, didFrom)
	if misesAccFrom == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "from account  %s not exists", didFrom)
	}
	if err != nil {
		return nil, err
	}
	misesAccTo, err := k.GetUserAccount(ctx, didTo)
	if misesAccTo == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "to account %s not exists", didTo)
	}
	if err != nil {
		return nil, err
	}

	if !k.HasUserRelationByMisesID(ctx, didFrom, didTo) {
		return nil, nil
	}
	rel := k.GetUserRelationByMisesID(ctx, didFrom, didTo)
	return &rel, nil
}

func (k *userMgr) userRelationFromBsonBytes(rawResult []byte) (*types.UserRelation, error) {
	var bsonVal bson.D
	err := bson.Unmarshal(rawResult, &bsonVal)
	if err != nil {
		return nil, err
	}
	var UserRelation = types.UserRelation{}
	if val, ok := bsonVal.Map()["uidfrom"]; ok {
		UserRelation.UidFrom = val.(string)
	}
	if val, ok := bsonVal.Map()["uidto"]; ok {
		UserRelation.UidTo = val.(string)
	}

	if val, ok := bsonVal.Map()["creator"]; ok {
		UserRelation.Creator = val.(string)
	}
	if val, ok := bsonVal.Map()["id"]; ok {
		UserRelation.Id = uint64(val.(int64))
	}
	if val, ok := bsonVal.Map()["isfollowing"]; ok {
		UserRelation.IsFollowing = val.(bool)
	}
	if val, ok := bsonVal.Map()["isblocking"]; ok {
		UserRelation.IsBlocking = val.(bool)
	}
	if val, ok := bsonVal.Map()["isreferredby"]; ok {
		UserRelation.IsReferredBy = val.(bool)
	}
	if val, ok := bsonVal.Map()["version"]; ok {
		UserRelation.Version = uint64(val.(int64))
	}
	return &UserRelation, nil
}

func (k *userMgr) GetUserRelations(ctx sdk.Context, relType uint64, didFrom string, lastDidTo string, limit int) ([]*types.UserRelation, error) {
	if didFrom == lastDidTo {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "from user must diff from last to user")
	}
	if limit > 100 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "limit must small than 100")
	}
	if limit < 1 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "limit must large than 1")
	}
	misesAccFrom, err := k.GetUserAccount(ctx, didFrom)
	if misesAccFrom == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "from account  %s not exists", didFrom)
	}
	if err != nil {
		return nil, err
	}
	if lastDidTo != "" {
		misesAccTo, err := k.GetUserAccount(ctx, lastDidTo)
		if misesAccTo == nil {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "last to account %s not exists", lastDidTo)
		}
		if err != nil {
			return nil, err
		}
	}
	if k.db == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "not implemented")
	}

	var UserRelations = []*types.UserRelation{}
	if db, ok := k.db.Raw().(*mongo.Database); ok {
		collection := db.Collection("UserRelation")
		filter := bson.M{
			"uidfrom": bson.M{"$eq": didFrom},
		}
		if lastDidTo != "" {
			filter["uidto"] = bson.M{"$gt": lastDidTo}
		}
		if relType&types.RelTypeBitFollow != 0 {
			filter["isfollowing"] = bson.M{"$eq": true}
		}

		if relType&types.RelTypeBitBlock != 0 {
			filter["isblocking"] = bson.M{"$eq": true}
		}

		if relType&types.RelTypeBitReferredBy != 0 {
			filter["isreferredby"] = bson.M{"$eq": true}
		}

		findOptions := options.Find()
		findOptions.SetSort(bson.M{"uidto": 1})

		cursor, err := collection.Find(context.Background(), filter, findOptions)
		if err != nil {
			return nil, err
		}
		for {
			if !cursor.Next(context.Background()) {
				break
			}

			rawResult := cursor.Current
			if err != nil {
				return nil, err
			}
			UserRelation, err := k.userRelationFromBsonBytes(rawResult)
			if err != nil {
				return nil, err
			}
			UserRelations = append(UserRelations, UserRelation)
			limit--
			if limit == 0 {
				break
			}
		}

	}

	return UserRelations, nil
}
