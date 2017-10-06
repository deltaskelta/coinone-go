package coinone

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// CoinBalance is part of the get balance response.
type CoinBalance struct {
	Avail   float64 `json:"avail,string"`
	Balance float64 `json:"balance,string"`
}

// NormalWallet is part of the get balance response.
type NormalWallet struct {
	Balance float64 `json:"balance,string"`
	Label   string  `json:"label"`
}

// BalanceResp is part of the get balance response.
type BalanceResp struct {
	Result        string         `json:"result"`
	ErrorCode     string         `json:"errorCode"`
	Btc           CoinBalance    `json:"btc"`
	Eth           CoinBalance    `json:"eth"`
	Etc           CoinBalance    `json:"etc"`
	Krw           CoinBalance    `json:"krw"`
	NormalWallets []NormalWallet `json:"normalWallets"`
}

// GetBalance calls the get balance endpoint.
func (c *API) GetBalance() (*BalanceResp, error) {
	URL := "https://api.coinone.co.kr/v2/account/balance/"

	payload := &Payload{}
	stringPayload, signature, err := c.GetAndSignPayload(payload)
	if err != nil {
		return nil, errors.Wrap(err, "coinone get balance make payload")
	}

	resp, err := c.Post(URL, stringPayload, signature)
	if err != nil {
		return nil, errors.Wrap(err, "coinone response")
	}

	var retResp BalanceResp
	err = json.NewDecoder(resp.Body).Decode(&retResp)
	if err != nil {
		return nil, errors.Wrap(err, "coinone get balance json decode")
	}

	if resp.StatusCode != http.StatusOK {
		return &retResp, errors.New("status code should be 200")
	}

	return &retResp, nil
}
