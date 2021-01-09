package keeper

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

// GetAuthCount get the total number of auth
func (k Keeper) GetAuthCount(ctx sdk.Context) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthCountKey))
	byteKey := types.KeyPrefix(types.AuthCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	count, err := strconv.ParseInt(string(bz), 10, 64)
	if err != nil {
		// Panic because the count should be always formattable to int64
		panic("cannot decode count")
	}

	return count
}

// SetAuthCount set the total number of auth
func (k Keeper) SetAuthCount(ctx sdk.Context, count int64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthCountKey))
	byteKey := types.KeyPrefix(types.AuthCountKey)
	bz := []byte(strconv.FormatInt(count, 10))
	store.Set(byteKey, bz)
}

func (k Keeper) CreateAuth(ctx sdk.Context, msg types.MsgCreateAuth) {
	// Create the auth
	count := k.GetAuthCount(ctx)
	var auth = types.Auth{
		Id:       strconv.FormatInt(count, 10),
		Identity: msg.Identity,
		Service:  msg.Creator,
		Key:      msg.Key,
		// Status:         msg.Status,
		// CreationDt:     msg.CreationDt,
		// AvailabilityDt: msg.AvailabilityDt,
	}

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	key := types.KeyPrefix(types.AuthKey + auth.Id)
	value := k.cdc.MustMarshalBinaryBare(&auth)
	store.Set(key, value)

	// Update auth count
	k.SetAuthCount(ctx, count+1)
}

func (k Keeper) UpdateAuth(ctx sdk.Context, auth types.Auth) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	b := k.cdc.MustMarshalBinaryBare(&auth)
	store.Set(types.KeyPrefix(types.AuthKey+auth.Id), b)
}

func (k Keeper) GetAuth(ctx sdk.Context, key string) types.Auth {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	var auth types.Auth
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.AuthKey+key)), &auth)
	return auth
}

func (k Keeper) HasAuth(ctx sdk.Context, id string) bool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	return store.Has(types.KeyPrefix(types.AuthKey + id))
}

func (k Keeper) GetAuthOwner(ctx sdk.Context, key string) string {
	return k.GetAuth(ctx, key).Service
}

// DeleteAuth deletes a auth
func (k Keeper) DeleteAuth(ctx sdk.Context, key string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	store.Delete(types.KeyPrefix(types.AuthKey + key))
}

func (k Keeper) GetAllAuth(ctx sdk.Context) (msgs []types.Auth) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.AuthKey))
	iterator := sdk.KVStorePrefixIterator(store, types.KeyPrefix(types.AuthKey))

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var msg types.Auth
		k.cdc.MustUnmarshalBinaryBare(iterator.Value(), &msg)
		msgs = append(msgs, msg)
	}

	return
}