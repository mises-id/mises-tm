package keeper

import (
	"bytes"
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/mises-id/mises-tm/x/misestm/types"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	addr, didtype, err := types.AddrFormDid(did)
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

	if db, ok := k.db.Raw().(*mongo.Database); ok {
		collection := db.Collection("UserRelation")
		filter := bson.M{
			"uidfrom": bson.M{"$eq": didFrom},
			"uidto":   bson.M{"$eq": didTo},
		}

		findOptions := options.FindOne()
		findOptions.SetSort(bson.M{"version": -1})

		result := collection.FindOne(context.Background(), filter, findOptions)
		if result.Err() != nil {
			if result.Err() == mongo.ErrNoDocuments {
				return nil, nil
			}
			return nil, result.Err()
		}
		rawResult, err := result.DecodeBytes()
		if err != nil {
			return nil, err
		}
		ret, err = k.userRelationFromBsonBytes(rawResult)
		if err != nil {
			return nil, err
		}
	}

	return ret, nil
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
	if val, ok := bsonVal.Map()["isFollowing"]; ok {
		UserRelation.IsFollowing = val.(bool)
	}
	if val, ok := bsonVal.Map()["isBlocking"]; ok {
		UserRelation.IsBlocking = val.(bool)
	}
	if val, ok := bsonVal.Map()["isReferredBy"]; ok {
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

	var UserRelations = []*types.UserRelation{}
	if db, ok := k.db.Raw().(*mongo.Database); ok {
		collection := db.Collection("UserRelation")
		filter := bson.M{
			"uidfrom":  bson.M{"$eq": didFrom},
			"isLatest": bson.M{"$eq": 1},
		}
		if lastDidTo != "" {
			filter["uidto"] = bson.M{"$gt": lastDidTo}
		}
		if relType&types.RelTypeBitFollow != 0 {
			filter["isFollowing"] = bson.M{"$eq": true}
		}

		if relType&types.RelTypeBitBlock != 0 {
			filter["isBlocking"] = bson.M{"$eq": true}
		}

		if relType&types.RelTypeBitReferredBy != 0 {
			filter["isReferredBy"] = bson.M{"$eq": true}
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

func (k *userMgr) setLatestUserRelation(rel *types.UserRelation) (err error) {
	filter := bson.M{
		"uidfrom": bson.M{"$eq": rel.UidFrom},
		"uidto":   bson.M{"$eq": rel.UidTo},
	}

	filter["version"] = bson.M{"$eq": rel.Version}

	err = k.setLatestTag(filter, 1)
	if err != nil {
		return err
	}

	if rel.Version > 0 {
		filter["version"] = bson.M{"$eq": rel.Version - 1}

		err = k.setLatestTag(filter, 0)
		if err != nil {
			return err
		}
	}

	return nil
}
func (k *userMgr) setLatestTag(filter bson.M, isLatest uint8) (err error) {
	bsonval := bson.M{"isLatest": isLatest}
	update := bson.M{"$set": bsonval}

	opts := &options.UpdateOptions{}
	opts.SetUpsert(false)
	if db, ok := k.db.Raw().(*mongo.Database); ok {
		collection := db.Collection("UserRelation")

		_, err = collection.UpdateOne(context.Background(), filter, update, opts)

	} else {
		err = nil
	}
	return err
}

func (k *userMgr) nodeKeyFromBsonBytes(rawResult []byte) ([]byte, error) {
	var bsonVal bson.D
	err := bson.Unmarshal(rawResult, &bsonVal)
	if err != nil {
		return nil, err
	}
	var nodeKey []byte
	if val, ok := bsonVal.Map()["node_key"]; ok {
		nodeKey = val.(primitive.Binary).Data
	}

	return nodeKey, nil
}

func (k *userMgr) OnWrite(key []byte, value []byte, delete bool) error {
	if delete {
		return nil
	}
	prefix := types.KeyPrefix("s/_/" + types.UserRelationKey)
	nodePrefix := types.KeyPrefix(types.UserRelationKey)
	nodeKey, _ := k.nodeKeyFromBsonBytes(value)

	if bytes.HasPrefix(key, prefix) || bytes.HasPrefix(nodeKey, nodePrefix) {
		ret, err := k.userRelationFromBsonBytes(value)
		if err != nil {
			return err
		}
		return k.setLatestUserRelation(ret)
	}
	return nil
}
