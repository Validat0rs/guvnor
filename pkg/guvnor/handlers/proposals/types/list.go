package types

import "time"

type Pagination struct {
	NextKey interface{} `json:"next_key"`
	Total   string      `json:"total"`
}

type TotalDeposit []struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type FinalTallyResult struct {
	Yes        string `json:"yes"`
	Abstain    string `json:"abstain"`
	No         string `json:"no"`
	NoWithVeto string `json:"no_with_veto"`
}

type Content struct {
	Type        string `json:"@type"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Recipient   string `json:"recipient"`
	Amount      Amount `json:"amount"`
}

type Amount []struct {
	Denom  string `json:"denom"`
	Amount string `json:"amount"`
}

type Proposals []struct {
	ProposalID       string           `json:"proposal_id"`
	Content          Content          `json:"content"`
	Status           string           `json:"status"`
	FinalTallyResult FinalTallyResult `json:"final_tally_result"`
	SubmitTime       time.Time        `json:"submit_time"`
	DepositEndTime   time.Time        `json:"deposit_end_time"`
	TotalDeposit     TotalDeposit     `json:"total_deposit"`
	VotingStartTime  time.Time        `json:"voting_start_time"`
	VotingEndTime    time.Time        `json:"voting_end_time"`
}

type List struct {
	Proposals  Proposals  `json:"proposals"`
	Pagination Pagination `json:"pagination"`
}
