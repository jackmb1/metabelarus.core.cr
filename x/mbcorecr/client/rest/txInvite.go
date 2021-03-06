package rest

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/metabelarus/mbcorecr/x/mbcorecr/types"
)

// Used to not have an error if strconv is unused
var _ = strconv.Itoa(42)

type createInviteRequest struct {
	BaseReq      rest.BaseReq `json:"base_req"`
	Inviter      string       `json:"inviter"`
	Level        string       `json:"level"`
	IdentityType string       `json:"identity_type"`
	TmpAddress   string       `json:"tmp_address"`
	TmpPubKey    string       `json:"tmp_pubkey"`
}

func createInviteHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createInviteRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Inviter)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		_, err = sdk.AccAddressFromBech32(req.TmpAddress)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		_, err = sdk.GetPubKeyFromBech32(
			sdk.Bech32PubKeyTypeAccPub, req.TmpPubKey,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		identityLevel, ok := types.IdentityLevel_value[req.Level]
		if !ok {
			rest.WriteErrorResponse(
				w, http.StatusBadRequest,
				fmt.Sprintf("Identity level: %s does not exist", req.Level),
			)
			return
		}

		identityType, ok := types.IdentityType_value[req.IdentityType]
		if !ok {
			rest.WriteErrorResponse(
				w, http.StatusBadRequest,
				fmt.Sprintf("Identity type: %s does not exist", req.IdentityType),
			)
			return
		}

		msg := types.NewMsgCreateInvite(
			req.Inviter,
			types.IdentityLevel(identityLevel),
			types.IdentityType(identityType),
			req.TmpAddress,
			req.TmpPubKey,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type acceptInviteRequest struct {
	BaseReq rest.BaseReq `json:"base_req"`
	Invite  string       `json:"invite"`
	Address string       `json:"address"`
	PubKey  string       `json:"pubkey"`
}

func acceptInviteHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req acceptInviteRequest
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
			return
		}

		baseReq := req.BaseReq.Sanitize()
		if !baseReq.ValidateBasic(w) {
			return
		}

		_, err := sdk.AccAddressFromBech32(req.Address)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		_, err = sdk.GetPubKeyFromBech32(
			sdk.Bech32PubKeyTypeAccPub, req.PubKey,
		)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		msg := types.NewMsgAcceptInvite(
			req.Invite,
			clientCtx.GetFromAddress().String(),
			req.Address,
			req.PubKey,
		)

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
