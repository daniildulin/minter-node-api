package responses

type Validator struct {
	PubKey string `json:"pubkey"`
	Signed bool   `json:"signed"`
}
