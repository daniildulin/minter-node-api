package minter_node_api

import (
	"encoding/json"
	"fmt"
	"github.com/daniildulin/minter-node-api/responses"
	"net/http"
	"time"
)

type MinterNodeApi struct {
	httpClient *http.Client
	link       string
}

func New(link string) *MinterNodeApi {
	return &MinterNodeApi{
		link:       link,
		httpClient: &http.Client{Timeout: 30 * time.Second},
	}
}

func (api *MinterNodeApi) SetLink(link string) {
	api.link = link
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

func (api *MinterNodeApi) GetCandidate(pubKey string) (*responses.CandidateResponse, error) {
	response := responses.CandidateResponse{}
	link := api.link + `/candidate?pubkey=` + pubKey
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

func (api *MinterNodeApi) getJson(url string, target interface{}) error {
	r, err := api.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
