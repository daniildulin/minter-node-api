package responses

type EstimateTxResponse struct {
	Response
	Result struct {
		Commission string `json:"commission"`
	} `json:"result"`
}
