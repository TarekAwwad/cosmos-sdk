package keeper

import (
	"context"
	"fmt"

	"cosmossdk.io/collections"
	"cosmossdk.io/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
	v1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
)

// AddVote adds a vote on a specific proposal
func (keeper Keeper) AddSecretVote(ctx context.Context, proposalID uint64, voterAddr sdk.AccAddress, cypherID string, metadata string) error {
	// Check if proposal is in voting period.
	inVotingPeriod, err := keeper.VotingPeriodProposals.Has(ctx, proposalID)
	if err != nil {
		return err
	}

	if !inVotingPeriod {
		return errors.Wrapf(types.ErrInactiveProposal, "%d", proposalID)
	}

	err = keeper.assertMetadataLength(metadata)
	if err != nil {
		return err
	}

	// TODO : check if cypherID is valid using ZK proof ?

	secretVote := v1.NewSecretVote(proposalID, voterAddr, cypherID, metadata)
	err = keeper.SecretVotes.Set(ctx, collections.Join(proposalID, voterAddr), secretVote)
	if err != nil {
		return err
	}

	// called after a vote on a proposal is cast
	err = keeper.Hooks().AfterProposalVote(ctx, proposalID, voterAddr)
	if err != nil {
		return err
	}

	sdkCtx := sdk.UnwrapSDKContext(ctx)
	sdkCtx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeProposalVote,
			sdk.NewAttribute(types.AttributeKeyVoter, voterAddr.String()),
			sdk.NewAttribute(types.AttributeKeyCypherID, cypherID),
			sdk.NewAttribute(types.AttributeKeyProposalID, fmt.Sprintf("%d", proposalID)),
		),
	)

	return nil
}
