package v1beta1

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// query endpoints supported by the governance Querier
const (
	QueryParams      = "params"
	QueryProposals   = "proposals"
	QueryProposal    = "proposal"
	QueryDeposits    = "deposits"
	QueryDeposit     = "deposit"
	QueryVotes       = "votes"
	QueryVote        = "vote"
	QuerySecretVote  = "secret_vote"
	QuerySecretVotes = "secret_votes"
	QueryTally       = "tally"

	ParamDeposit  = "deposit"
	ParamVoting   = "voting"
	ParamTallying = "tallying"
)

// QueryProposalParams Params for queries:
// - 'custom/gov/proposal'
// - 'custom/gov/deposits'
// - 'custom/gov/tally'
type QueryProposalParams struct {
	ProposalID uint64
}

// NewQueryProposalParams creates a new instance of QueryProposalParams
func NewQueryProposalParams(proposalID uint64) QueryProposalParams {
	return QueryProposalParams{
		ProposalID: proposalID,
	}
}

// QueryProposalVotesParams used for queries to 'custom/gov/votes'.
type QueryProposalVotesParams struct {
	ProposalID uint64
	Page       int
	Limit      int
}

// NewQueryProposalVotesParams creates new instance of the QueryProposalVotesParams.
func NewQueryProposalVotesParams(proposalID uint64, page, limit int) QueryProposalVotesParams {
	return QueryProposalVotesParams{
		ProposalID: proposalID,
		Page:       page,
		Limit:      limit,
	}
}

// QueryProposalSecretVotesParams used to query 'custom/gov/secretvotes'.
type QueryProposalSecretVotesParams struct {
	ProposalID uint64
	Page       int
	Limit      int
}

// NewQueryProposalSecretVotesParams creates new instance of the QueryProposalSecretVotesParams.
func NewQueryProposalSecretVotesParams(proposalID uint64, page, limit int) QueryProposalSecretVotesParams {
	return QueryProposalSecretVotesParams{
		ProposalID: proposalID,
		Page:       page,
		Limit:      limit,
	}
}

// QueryDepositParams params for query 'custom/gov/deposit'
type QueryDepositParams struct {
	ProposalID uint64
	Depositor  sdk.AccAddress
}

// NewQueryDepositParams creates a new instance of QueryDepositParams
func NewQueryDepositParams(proposalID uint64, depositor sdk.AccAddress) QueryDepositParams {
	return QueryDepositParams{
		ProposalID: proposalID,
		Depositor:  depositor,
	}
}

// QueryVoteParams Params for query 'custom/gov/vote'
type QueryVoteParams struct {
	ProposalID uint64
	Voter      sdk.AccAddress
}

// NewQueryVoteParams creates a new instance of QueryVoteParams
func NewQueryVoteParams(proposalID uint64, voter sdk.AccAddress) QueryVoteParams {
	return QueryVoteParams{
		ProposalID: proposalID,
		Voter:      voter,
	}
}

// QuerySecretVoteParams Params for query 'custom/gov/secretvote'
type QuerySecretVoteParams struct {
	ProposalID uint64
	Voter      sdk.AccAddress
}

// NewQuerySecretVoteParams creates a new instance of QuerySecretVoteParams
func NewQuerySecretVoteParams(proposalID uint64, voter sdk.AccAddress) QuerySecretVoteParams {
	return QuerySecretVoteParams{
		ProposalID: proposalID,
		Voter:      voter,
	}
}

// QueryProposalsParams Params for query 'custom/gov/proposals'
type QueryProposalsParams struct {
	Page           int
	Limit          int
	Voter          sdk.AccAddress
	Depositor      sdk.AccAddress
	ProposalStatus ProposalStatus
}

// NewQueryProposalsParams creates a new instance of QueryProposalsParams
func NewQueryProposalsParams(page, limit int, status ProposalStatus, voter, depositor sdk.AccAddress) QueryProposalsParams {
	return QueryProposalsParams{
		Page:           page,
		Limit:          limit,
		Voter:          voter,
		Depositor:      depositor,
		ProposalStatus: status,
	}
}
