package minter_node_api

import (
	"encoding/json"
	"github.com/daniildulin/minter-node-api/responses"
	"net/http"
)

type MinterNodeApi struct {
	httpClient *http.Client
	link       string
}

func New(httpClient *http.Client, link string) *MinterNodeApi {
	return &MinterNodeApi{
		httpClient: httpClient,
		link:       link,
	}
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

func (api *MinterNodeApi) getJson(url string, target interface{}) error {
	r, err := api.httpClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
