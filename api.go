package minter_node_api

import (
	"encoding/json"
	"fmt"
	"github.com/daniildulin/minter-node-api/responses"
	"github.com/valyala/fasthttp"
	"strconv"
	"strings"
)

type MinterNodeApi struct {
	link string
}

func New(link string) *MinterNodeApi {
	return &MinterNodeApi{
		link: link,
	}
}

func (api *MinterNodeApi) SetLink(link string) {
	api.link = link
}

func (api *MinterNodeApi) GetLink() string {
	return api.link
}

func (api *MinterNodeApi) GetStatus() (*responses.StatusResponse, error) {
	response := responses.StatusResponse{}
	link := api.link + `/status`
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetBlock(height uint64) (*responses.BlockResponse, error) {
	response := responses.BlockResponse{}
	link := api.link + `/block?height=` + fmt.Sprint(height)
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetBlockEvents(height uint64) (*responses.EventsResponse, error) {
	response := responses.EventsResponse{}
	link := api.link + `/events?height=` + fmt.Sprint(height)
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetBlockValidators(height uint64) (*responses.ValidatorsResponse, error) {
	response := responses.ValidatorsResponse{}
	link := api.link + `/validators?height=` + fmt.Sprint(height)
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetCandidate(pubKey string, height uint64) (*responses.CandidateResponse, error) {
	response := responses.CandidateResponse{}
	link := api.link + `/candidate?pubkey=` + pubKey + `&height=` + strconv.Itoa(int(height))
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetCandidates(height uint64) (*responses.BlockCandidatesResponse, error) {
	response := responses.BlockCandidatesResponse{}
	link := api.link + `/candidates?height=` + strconv.Itoa(int(height))
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetCoinInfo(symbol string) (*responses.CoinInfoResponse, error) {
	response := responses.CoinInfoResponse{}
	link := api.link + `/coin_info?symbol=` + symbol
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetAddress(address string) (*responses.AddressResponse, error) {
	response := responses.AddressResponse{}
	link := api.link + `/address?address=` + strings.Title(address)
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetAddresses(addresses []string, height uint64) (*responses.BalancesResponse, error) {
	response := responses.BalancesResponse{}
	queryStr := "[" + strings.Join(addresses, ",") + "]"
	link := api.link + `/addresses?addresses=` + queryStr + `&height=` + strconv.Itoa(int(height))
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetEstimateTx(tx string) (*responses.EstimateTxResponse, error) {
	response := responses.EstimateTxResponse{}
	link := api.link + `/estimate_tx_commission?tx=` + tx
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetEstimateCoinBuy(coinToSell string, coinToBuy string, value string) (*responses.EstimateCoinBuyResponse, error) {
	response := responses.EstimateCoinBuyResponse{}
	link := api.link + `/estimate_coin_buy?coin_to_sell=` + coinToSell + `&coin_to_buy=` + coinToBuy + `&value_to_buy=` + value
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetEstimateCoinSell(coinToSell string, coinToBuy string, value string) (*responses.EstimateCoinSellResponse, error) {
	response := responses.EstimateCoinSellResponse{}
	link := api.link + `/estimate_coin_sell?coin_to_sell=` + coinToSell + `&coin_to_buy=` + coinToBuy + `&value_to_sell=` + value
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) GetMinGasPrice() (*responses.GasResponse, error) {
	response := responses.GasResponse{}
	link := api.link + `/min_gas_price`
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) PushTransaction(tx string) (*responses.SendTransactionResponse, error) {
	response := responses.SendTransactionResponse{}
	link := api.link + `/send_transaction?tx=0x` + tx
	err := api.getJson(link, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
}

func (api *MinterNodeApi) getJson(url string, target interface{}) error {
	_, body, err := fasthttp.Get(nil, url)
	if err != nil {
		return err
	}
	return json.Unmarshal(body, target)
}
