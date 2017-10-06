package coinone

import (
	"encoding/json"
	"net/http"

	"github.com/pkg/errors"
)

// CancelOrderResp is the response that is returned when cancelling a single order
type CancelOrderResp struct {
	Result    string `json:"result"`
	ErrorCode int    `json:"errorCode,string"`
}

// CancelOrder is the function that calls the cancel order endpoint.
func (c *API) CancelOrder(currency, orderID string, qty float64, isAsk int, price int64) (*CancelOrderResp, error) {
	URL := "https://api.coinone.co.kr/v2/order/cancel/"

	payload := &Payload{
		OrderID:  orderID,
		Price:    price,
		Qty:      qty,
		IsAsk:    isAsk,
		Currency: currency,
	}

	stringPayload, signature, err := c.GetAndSignPayload(payload)
	if err != nil {
		return nil, errors.Wrap(err, "Cancel all orders get payload")
	}

	resp, err := c.Post(URL, stringPayload, signature)
	if err != nil {
		return nil, errors.Wrap(err, "cancel all orders post")
	}

	var retResp CancelOrderResp
	err = json.NewDecoder(resp.Body).Decode(&retResp)
	if err != nil {
		return nil, errors.Wrap(err, "cancel all orders decode json")
	}

	if resp.StatusCode != http.StatusOK {
		return &retResp, errors.New("status should be 200")
	}

	return &retResp, nil
}

// OrderResp is the struct that is returned from making an order
type OrderResp struct {
	Result    string `json:"result"`
	OrderID   string `json:"orderId"`
	ErrorCode int    `json:"errorCode,string"`
}

// LimitBuy calls the limit buy endpoint.
func (c *API) LimitBuy(currency string, price int64, qty float64) (*OrderResp, error) {
	URL := "https://api.coinone.co.kr/v2/order/limit_buy/"

	payload := &Payload{
		Price:    price,
		Qty:      qty,
		Currency: currency,
	}

	stringPayload, signature, err := c.GetAndSignPayload(payload)
	if err != nil {
		return nil, errors.Wrap(err, "coinone limit buy")
	}

	resp, err := c.Post(URL, stringPayload, signature)
	if err != nil {
		return nil, errors.Wrap(err, "coinone limit buy post")
	}

	var retResp OrderResp
	err = json.NewDecoder(resp.Body).Decode(&retResp)
	if err != nil {
		return nil, errors.Wrap(err, "coinone limit buy json decode")
	}

	if resp.StatusCode != http.StatusOK {
		return &retResp, errors.New("status should be 200")
	}

	return &retResp, nil
}

// LimitSell calls the limit sell endpoint.
func (c *API) LimitSell(currency string, price int64, qty float64) (*OrderResp, error) {
	URL := "https://api.coinone.co.kr/v2/order/limit_sell/"

	payload := &Payload{
		Price:    price,
		Qty:      qty,
		Currency: currency,
	}

	stringPayload, signature, err := c.GetAndSignPayload(payload)
	if err != nil {
		return nil, errors.Wrap(err, "coinone limit sell get payload")
	}

	resp, err := c.Post(URL, stringPayload, signature)
	if err != nil {
		return nil, errors.Wrap(err, "coinone limit sell post")
	}

	var retResp OrderResp
	err = json.NewDecoder(resp.Body).Decode(&retResp)
	if err != nil {
		return nil, errors.Wrap(err, "coinone limit sell json decode")
	}

	if resp.StatusCode != http.StatusOK {
		return &retResp, errors.New("status should be 200")
	}

	return &retResp, nil
}

// LimitOrderResp is the response that is returned from listing orders.
type LimitOrderResp struct {
	Result      string  `json:"result"`
	ErrorCode   int     `json:"errorCode,string"`
	LimitOrders []Order `json:"limitOrders"`
}

// Order is the actual order that is returned in an order list.
type Order struct {
	Index     int     `json:"index,string"`
	Timestamp int64   `json:"timestamp,string"`
	Price     int64   `json:"price,string"`
	Qty       float64 `json:"qty,string"`
	OrderID   string  `json:"orderId"`
	Type      string  `json:"type"`
	FeeRate   float64 `json:"feeRate,string"`
	Fee       float64 `json:"fee,string"`
}

// LimitOrders calls the list orders endpoint.
func (c *API) LimitOrders(coin string) (*LimitOrderResp, error) {
	URL := "https://api.coinone.co.kr/v2/order/limit_orders/"

	payload := &Payload{
		Currency: coin,
	}

	stringPayload, signature, err := c.GetAndSignPayload(payload)
	if err != nil {
		return nil, errors.Wrap(err, "coinone list orders get payload")
	}

	resp, err := c.Post(URL, stringPayload, signature)
	if err != nil {
		return nil, errors.Wrap(err, "coinone list orders post")
	}

	var retResp LimitOrderResp
	err = json.NewDecoder(resp.Body).Decode(&retResp)
	if err != nil {
		return nil, errors.Wrap(err, "coinone list orders json decode")
	}

	if resp.StatusCode != http.StatusOK {
		return &retResp, errors.New("status should be 200")
	}

	return &retResp, nil
}

// CompleteOrderResp is the response that is given when querying complete orders
type CompleteOrderResp struct {
	Result         string  `json:"result"`
	ErrorCode      int     `json:"errorCode,string"`
	CompleteOrders []Order `json:"completeOrders"`
}

// CompleteOrders call the complete orders endpoint.
func (c *API) CompleteOrders(coin string) (*CompleteOrderResp, error) {
	URL := "https://api.coinone.co.kr/v2/order/complete_orders/"

	payload := &Payload{
		Currency: coin,
	}

	stringPayload, signature, err := c.GetAndSignPayload(payload)
	if err != nil {
		return nil, errors.Wrap(err, "coinone complete orders get payload")
	}

	resp, err := c.Post(URL, stringPayload, signature)
	if err != nil {
		return nil, errors.Wrap(err, "coinone complete orders post")
	}

	var retResp CompleteOrderResp
	err = json.NewDecoder(resp.Body).Decode(&retResp)
	if err != nil {
		return nil, errors.Wrap(err, "coinone complete orders json decode")
	}

	if resp.StatusCode != http.StatusOK {
		return &retResp, errors.New("status should be 200")
	}

	return &retResp, nil
}
