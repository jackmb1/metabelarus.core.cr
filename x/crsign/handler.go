package crsign

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/metabelarus/mbcorecr/x/crsign/keeper"
	"github.com/metabelarus/mbcorecr/x/crsign/types"
)

// NewHandler ...
func NewHandler(k keeper.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())

		switch msg := msg.(type) {
		// this line is used by starport scaffolding # 1
		case *types.MsgRequestAuth:
			return handleMsgRequestAuth(ctx, k, msg)

		case *types.MsgConfirmAuth:
			return handleMsgConfirmAuth(ctx, k, msg)

		case *types.MsgCreateRecord:
			return handleMsgCreateRecord(ctx, k, msg)

		case *types.MsgUpdateRecord:
			return handleMsgUpdateRecord(ctx, k, msg)

		case *types.MsgDeleteRecord:
			return handleMsgDeleteRecord(ctx, k, msg)

		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}
