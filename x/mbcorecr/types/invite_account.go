package types

import (
	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

type InviteAccount struct {
	Address string
	PubKey  string
	ctx     *sdk.Context
}

// NewInviteAccount - Generate payload object for Invite and
// Identity Account responses.
func NewInviteAccount(address string, pubKey string, ctx *sdk.Context, authKeeper AccountKeeper) (*InviteAccount, error) {
	inviteAcc := &InviteAccount{
		Address: address,
		PubKey:  pubKey,
		ctx:     ctx,
	}

	addr, err := inviteAcc.GetAccAddress()
	if err != nil {
		return nil, err
	}

	newAcc := authKeeper.NewAccountWithAddress(*ctx, addr)
	if err := newAcc.SetPubKey(inviteAcc.GetAccPubKey()); err != nil {
		return nil, sdkerrors.Wrap(ErrNewAccount, err.Error())
	}
	authKeeper.SetAccount(*ctx, newAcc)

	return inviteAcc, nil
}

func (this *InviteAccount) GetAccAddress() (sdk.AccAddress, error) {
	inviteAddr, err := sdk.AccAddressFromBech32(this.Address)
	if err != nil {
		return nil, sdkerrors.Wrap(ErrNewAccount, err.Error())
	}

	return inviteAddr, nil
}

func (this *InviteAccount) GetAccPubKey() crypto.PubKey {
	return sdk.MustGetPubKeyFromBech32(sdk.Bech32PubKeyTypeAccPub, this.PubKey)
}

func (this *InviteAccount) SetBalances(bank BankKeeper, balances sdk.Coins) error {
	inviteAddr, err := this.GetAccAddress()
	if err != nil {
		return err
	}

	if err := bank.SetBalances(
		*this.ctx,
		inviteAddr,
		balances,
	); err != nil {
		return sdkerrors.Wrap(ErrNewAccount, err.Error())
	}

	return nil
}
