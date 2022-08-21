package coinranking

import (
	"encoding/json"
	"io"
	"net/http"
)

const URL_API = "https://api.coinranking.com/v2"

type Client struct {
	http.Client
	reqHeader http.Header
}

func NewClient(token string) *Client {
	header := http.Header{"access_token": {token}}
	return &Client{reqHeader: header}
}

func (c *Client) GetCoin(uuid string) (*Coin, error) {
	var jsonResponse Response

	req, err := http.NewRequest("GET", URL_API+"/coin/"+uuid, nil)

	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)

	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(res.Body)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(data, &jsonResponse); err != nil {
		return nil, err
	}

	return jsonResponse.Data.Coin, jsonResponse.Check()
}
