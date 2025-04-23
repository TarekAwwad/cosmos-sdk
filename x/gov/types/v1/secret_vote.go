package v1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewVote creates a new Vote instance
func NewSecretVote(proposalID uint64, voter sdk.AccAddress, cypherID string, metadata string) SecretVote {
	return SecretVote{ProposalId: proposalID, Voter: voter.String(), CypherId: cypherID, Metadata: metadata}
}

// Empty returns whether a vote is empty.
func (v SecretVote) Empty() bool {
	return v.ProposalId == 0 || v.Voter == "" || len(v.CypherId) == 0
}

// Votes is a collection of Vote objects
type SecretVotes []*SecretVote
