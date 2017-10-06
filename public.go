package coinone

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
)

// Quote is the basic quote that coinone returns.
type Quote struct {
	Volume   float64 `json:"volume,string"`
	Last     float64 `json:"last,string"`
	High     float64 `json:"high,string"`
	Currency string  `json:"currency"`
	Low      float64 `json:"low,string"`
	First    float64 `json:"first,string"`
}

// Response is the response that is returned from coinone.
type Response struct {
	Timestamp string `json:"timestamp"`
	ErrorCode string `json:"errorCode"`
	Result    string `json:"result"`
	Etc       Quote  `json:"etc"`
	Btc       Quote  `json:"btc"`
	Eth       Quote  `json:"eth"`
	Xrp       Quote  `json:"xrp"`
}

// GetCurrencyQuote gets all the currency quotes for coinone and puts them into one
// object.
func (c *API) GetCurrencyQuote() (*Response, error) {
	URL := "https://api.coinone.co.kr/ticker/?currency=all"

	// this first cycle does not have the current bid and ask price that is in the order
	// book so I have to get that later
	resp, err := c.Client.Get(URL)
	if err != nil {
		return nil, errors.Wrap(err, "coinone request")
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading coinone resp")
	}

	coinOneBuffer := bytes.NewBuffer(respBytes)

	var coinoneResp Response
	err = json.NewDecoder(coinOneBuffer).Decode(&coinoneResp)
	if err != nil {
		return nil, errors.Wrap(err, "decoding coinone resp")
	}

	if resp.StatusCode != http.StatusOK {
		return &coinoneResp, errors.New("non 200 status code")
	}

	return &coinoneResp, nil
}

// PriceQuantity is the price and quantity number returned from coinone.
type PriceQuantity struct {
	Price float64 `json:"price,string"`
	Qty   float64 `json:"qty,string"`
}

// Orderbook gives the bid and ask prices which come in a different place than
// the other prices.
type Orderbook struct {
	Timestamp int64           `json:"timestamp,string"`
	ErrorCode int             `json:"errorCode,string"`
	Currency  string          `json:"currency"`
	Result    string          `json:"result"`
	Bid       []PriceQuantity `json:"bid"`
	Ask       []PriceQuantity `json:"ask"`
}

// GetOrderbookQuote gets the orderbook for a certain coin which contains the orders and
// the bids and asks for all of them.
func (c *API) GetOrderbookQuote(currency string) (*Orderbook, error) {
	// this orderbook is going to be where the bif and ask prices for coinone are.
	url := fmt.Sprintf("https://api.coinone.co.kr/orderbook/?currency=%s", currency)
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "coinone get bid/ask at %s", currency)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "reading coinone bid/ask resp at %s", currency)
	}

	coinoneBidAskBuffer := bytes.NewBuffer(respBytes)

	var coinoneOrderbookResp Orderbook
	err = json.NewDecoder(coinoneBidAskBuffer).Decode(&coinoneOrderbookResp)
	if err != nil {
		return nil, errors.Wrapf(err, "coinone bid/ask decode at %s", currency)
	}

	if resp.StatusCode != http.StatusOK {
		return &coinoneOrderbookResp, errors.New("non 200 status code")
	}

	return &coinoneOrderbookResp, nil
}

// BidAskQuote is just like the orderbook quote but the bids and asks are taken sing
// and sorted.
type BidAskQuote struct {
	Timestamp int64
	ErrorCode int
	Currency  string
	Result    string
	Bid       PriceQuantity
	Ask       PriceQuantity
}

// GetBidAskQuote is like the orderbook but it takes care of the work of sorting out the
// bids and asks to get the highest bid and the lowest ask.
func (c *API) GetBidAskQuote(currency string) (*BidAskQuote, error) {
	url := fmt.Sprintf("https://api.coinone.co.kr/orderbook/?currency=%s", currency)
	resp, err := c.Client.Get(url)
	if err != nil {
		return nil, errors.Wrapf(err, "coinone get bid/ask at %s", currency)
	}
	defer resp.Body.Close()

	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrapf(err, "reading coinone bid/ask resp at %s", currency)
	}

	coinoneBidAskBuffer := bytes.NewBuffer(respBytes)

	var coinoneOrderbookResp Orderbook
	err = json.NewDecoder(coinoneBidAskBuffer).Decode(&coinoneOrderbookResp)
	if err != nil {
		return nil, errors.Wrapf(err, "coinone bid/ask decode at %s", currency)
	}

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("non 200 status code")
	}

	bidAskQuote := BidAskQuote{
		Currency:  coinoneOrderbookResp.Currency,
		ErrorCode: coinoneOrderbookResp.ErrorCode,
		Result:    coinoneOrderbookResp.Result,
		Timestamp: coinoneOrderbookResp.Timestamp,
		Ask:       coinoneOrderbookResp.Ask[0],
		Bid:       coinoneOrderbookResp.Bid[0],
	}

	for _, v := range coinoneOrderbookResp.Ask {
		if v.Price < bidAskQuote.Ask.Price {
			bidAskQuote.Ask = v
		}
	}

	for _, v := range coinoneOrderbookResp.Bid {
		if v.Price > bidAskQuote.Bid.Price {
			bidAskQuote.Bid = v
		}
	}

	return &bidAskQuote, nil

}
