package responses

type CandidateResponse struct {
	Response
	Result struct {
		CandidateAddress string  `json:"candidate_address"`
		TotalStake       string  `json:"total_stake"`
		PubKey           string  `json:"pubkey"`
		Commission       string  `json:"commission"`
		CreatedAtBlock   string  `json:"created_at_block"`
		Status           byte    `json:"status"`
		Stakes           []Stake `json:"stakes"`
	} `json:"result"`
}

type BlockCandidatesResponse struct {
	Response
	Result []struct {
		RewardAddress  string `json:"reward_address"`
		OwnerAddress   string `json:"owner_address"`
		TotalStake     string `json:"total_stake"`
		PubKey         string `json:"pubkey"`
		Commission     string `json:"commission"`
		CreatedAtBlock string `json:"created_at_block"`
		Status         byte   `json:"status"`
	} `json:"result"`
}

type Stake struct {
	Owner    string `json:"owner"`
	Coin     string `json:"coin"`
	Value    string `json:"value"`
	BipValue string `json:"bip_value"`
}
