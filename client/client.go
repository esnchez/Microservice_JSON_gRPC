package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/esnchez/types"
)

//new folder because this needs to be importable and may be used in other microservices.
//making the client for our service

type Client struct {
	endpoint string
}

func New(endpoint string) *Client{
	return &Client{
		endpoint: endpoint,
	}
}

func (s *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error){
	endpoint := fmt.Sprintf("%s?ticker=%s", s.endpoint, ticker )

	req, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return nil,err
	}

	//IF we are are receveing a bad response from the api, managing the error we output
	if resp.StatusCode != http.StatusOK {
		httpErr := map[string]any{}
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Service responsed with non Ok status code: %s", httpErr["error"])
	}

	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}
	return priceResp, nil;

} 