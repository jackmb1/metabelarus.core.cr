package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

func (k Keeper) GetServicesFromId(ctx sdk.Context, id string) types.Id2Service {
	var id2services types.Id2Service
	k.cdc.MustUnmarshalBinaryBare(
		prefix.NewStore(
			ctx.KVStore(k.storeKey),
			types.KeyPrefix(types.Id2ServicesKey),
		).Get(
			types.KeyPrefix(types.Id2ServicesKey+id),
		),
		&id2services,
	)

	return id2services
}

func (k Keeper) AddService2Id(ctx sdk.Context, id string, service string) {
	id2services := k.GetServicesFromId(ctx, id)
	for _, srv := range id2services.Services {
		if srv == service {
			return
		}
	}

	if id2services.Services == nil || len(id2services.Services) == 0 {
		id2services.Services = []string{service}
	} else {
		id2services.Services = append(id2services.Services, service)
	}

	prefix.NewStore(
		ctx.KVStore(k.storeKey), types.KeyPrefix(types.Id2ServicesKey),
	).Set(
		types.KeyPrefix(types.Id2ServicesKey+id),
		k.cdc.MustMarshalBinaryBare(&id2services),
	)
}
