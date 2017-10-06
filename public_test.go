package coinone

import (
	"errors"
	"testing"
)

func makeAPI() (*API, error) {

	if apiKey == "" || secretKey == "" {
		return nil, errors.New("You need to define an API key and a secret key somewhere" +
			"in the package")
	}

	coinoneAPI, err := NewAPI(apiKey, secretKey)
	if err != nil {
		return nil, err
	}

	return coinoneAPI, nil
}

func TestGetQuote(t *testing.T) {
	api, err := makeAPI()
	if err != nil {
		t.Error(err)
	}

	quote, err := api.GetCurrencyQuote()
	if err != nil {
		t.Error(err)
	}

	if quote == nil {
		t.Error("quote should not be nil")
	}
}

type test struct {
	Coin string
}

func TestGetOrderBook(t *testing.T) {
	api, err := makeAPI()
	if err != nil {
		t.Error(err)
	}

	tests := []test{
		{Coin: BTC},
		{Coin: ETH},
		{Coin: ETC},
		{Coin: XRP},
	}

	for _, v := range tests {
		orderbook, err := api.GetOrderbookQuote(v.Coin)
		if err != nil {
			t.Error(v.Coin)
			t.Error(err)
		}

		if orderbook == nil {
			t.Error("orderbook should not be nil")
		}
	}
}
