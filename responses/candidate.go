package responses

type CandidateResponse struct {
	Response
	Result struct {
		CandidateAddress string  `json:"candidate_address"`
		TotalStake       string  `json:"total_stake"`
		Pubkey           string  `json:"pubkey"`
		Commission       string  `json:"commission"`
		CreatedAtBlock   string  `json:"created_at_block"`
		Status           byte    `json:"status"`
		Stakes           []Stake `json:"stakes"`
	} `json:"result"`
}

type Stake struct {
	Owner    string `json:"owner"`
	Coin     string `json:"coin"`
	Value    string `json:"value"`
	BipValue string `json:"bip_value"`
}
