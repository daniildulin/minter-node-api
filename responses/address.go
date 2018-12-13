package responses

type AddressResponse struct {
	Response
	Result struct {
		TransactionCount string            `json:"transaction_count"`
		Balance          map[string]string `json:"balance"`
	} `json:"result"`
}
