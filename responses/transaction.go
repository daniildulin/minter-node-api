package responses

type TransactionResponse struct {
	Response
	Result Transaction `json:"result"`
}

type Transaction struct {
	Hash        string                  `json:"hash"`
	From        string                  `json:"from"`
	Type        uint8                   `json:"type"`
	Nonce       string                  `json:"nonce"`
	GasPrice    string                  `json:"gas_price"`
	GasCoin     string                  `json:"gas_coin"`
	GasUsed     string                  `json:"gas_used"`
	Gas         string                  `json:"gas"`
	Payload     string                  `json:"payload"`
	ServiceData string                  `json:"service_data"`
	RawTx       string                  `json:"raw_tx"`
	Log         *string                 `json:"log"`
	Data        *map[string]interface{} `json:"data"`
	Tags        *map[string]string      `json:"tags"`
}

type TransactionData struct {
	Coin                 *string              `json:"coin"`
	To                   *string              `json:"to"`
	Value                *string              `json:"value"`
	CoinToSell           *string              `json:"coin_to_sell"`
	CoinToBuy            *string              `json:"coin_to_buy"`
	ValueToSell          *string              `json:"value_to_sell"`
	ValueToBuy           *string              `json:"value_to_buy"`
	Name                 *string              `json:"name"`
	Symbol               *string              `json:"symbol"`
	InitialAmount        *string              `json:"initial_amount"`
	InitialReserve       *string              `json:"initial_reserve"`
	ConstantReserveRatio *string              `json:"constant_reserve_ratio"`
	Address              *string              `json:"address"`
	PubKey               *string              `json:"pub_key"`
	Commission           *string              `json:"commission"`
	Stake                *string              `json:"stake"`
	Proof                *string              `json:"proof"`
	RawCheck             *[]byte              `json:"raw_check"`
	ToCoinSymbol         *string              `json:"to_coin_symbol"`
	FromCoinSymbol       *string              `json:"from_coin_symbol"`
	Threshold            *string              `json:"threshold"`
	RewardAddress        *string              `json:"reward_address"`
	OwnerAddress         *string              `json:"owner_address"`
	ReceiversList        *[]MultiSendReceiver `json:"list"`
}

type MultiSendReceiver struct {
	Coin  string `json:"coin"`
	To    string `json:"to"`
	Value string `json:"value"`
}

type SendTransactionResponse struct {
	Response
	Result *SendTransactionResult `json:"result"`
}

type SendTransactionResult struct {
	Code int32  `json:"code"`
	Data string `json:"data"`
	Log  string `json:"log"`
	Hash string `json:"hash"`
}
