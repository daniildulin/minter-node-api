package responses

import "time"

type BlockResponse struct {
	Response
	Result struct {
		Hash         string        `json:"hash"`
		Height       string        `json:"height"`
		Time         time.Time     `json:"time"`
		TxCount      string        `json:"num_txs"`
		TotalTx      string        `json:"total_txs"`
		BlockReward  string        `json:"block_reward"`
		Size         string        `json:"size"`
		Proposer     string        `json:"proposer"`
		Transactions []Transaction `json:"transactions"`
		Validators   []Validator   `json:"validators"`
	} `json:"result"`
}
