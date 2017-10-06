package coinone

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
)

// BTC and other variables here are the coin names for the coinone site.
var (
	BTC = "btc"
	ETH = "eth"
	ETC = "etc"
	XRP = "xrp"
)

// API is the basic API structure for coinone.
type API struct {
	APIKey       string
	SecretKey    string
	RefreshToken string
	Nonce        int64
	Client       http.Client
}

/*
	PAYLOAD IS: acces_token, nonce and other data needed for request, base64 encoded
				X-COINONE-PAYLOAD
	SIGNATURE IS: Payload hmac with secret key sha512
				X-COINONE-SIGNATURE
*/

// NewAPI takes in an API and secret string and returns an API object to be used.
func NewAPI(APIKey, SecretKey string) (*API, error) {
	if APIKey == "" || SecretKey == "" {
		return nil, errors.New("apikey and secret key cannot be empty")
	}

	c := API{
		APIKey:    APIKey,
		SecretKey: SecretKey,
		Nonce:     time.Now().Unix(),
		Client:    http.Client{},
	}

	return &c, nil
}

// GetNonce sets the nonce one integer value higher and then returns it for use in a
// payload.
func (c *API) GetNonce() int64 {
	c.Nonce++
	return c.Nonce
}

// Payload is the basic payload that will be sent to API endpoints, not all of the
// fields are to be used in every endpoint.
type Payload struct {
	AccessToken string  `json:"access_token"`
	Nonce       int64   `json:"nonce"`
	Currency    string  `json:"currency"`
	OrderID     string  `json:"order_id"`
	Price       int64   `json:"price,string"`
	Qty         float64 `json:"qty,string"`
	IsAsk       int     `json:"is_ask,string"`
	Address     string  `json:"address"`
	AuthNumber  int     `json:"auth_number,string"`
}

// GetAndSignPayload is used in the endpoint functions to easily create a payload and sign
// everything.
func (c *API) GetAndSignPayload(payload *Payload) (retPayload, signature string, err error) {
	payload.Nonce = c.GetNonce()
	payload.AccessToken = c.APIKey

	jsonPayload, err := json.Marshal(payload)
	if err != nil {
		return "", "", errors.Wrap(err, "json marshal coinone payload")
	}

	b64Payload := base64.StdEncoding.EncodeToString(jsonPayload)

	secretUpper := strings.ToUpper(c.SecretKey)
	hash := hmac.New(sha512.New, []byte(secretUpper))
	hash.Write([]byte(b64Payload))
	sig := fmt.Sprintf("%064x", hash.Sum(nil))

	return b64Payload, sig, nil
}

// Post creates a post request and executes it returning the response.
func (c *API) Post(URL, payload, signature string) (*http.Response, error) {
	body := bytes.NewBuffer([]byte(payload))

	req, err := http.NewRequest("POST", URL, body)
	if err != nil {
		return nil, errors.Wrap(err, "coinone make request")
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-COINONE-PAYLOAD", payload)
	req.Header.Set("X-COINONE-SIGNATURE", signature)

	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "coinone post")
	}

	return resp, nil
}
