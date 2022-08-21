package coinranking

import "errors"

type Response struct {
	Status string `json:"status"`
	Data   struct {
		Coin *Coin `json:"coin"`
	} `json:"data"`
	Code    string `json:"code"`
	Message string `json:"message"`
}

func (r Response) Check() error {
	if r.Status == "success" {
		return nil
	}

	return errors.New(r.Code + ":" + r.Message)
}
