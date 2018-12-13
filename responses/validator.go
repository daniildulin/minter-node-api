package responses

type ValidatorsResponse struct {
	Response
	Signed []Validator `json:"result"`
}

type Validator struct {
	PubKey      string  `json:"pubkey"`
	Signed      bool    `json:"signed"`
	VotingPower *string `json:"voting_power"`
}
