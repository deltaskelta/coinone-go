package coinone

import (
	"testing"
)

type CancelOrderTest struct {
	Currency string
	OrderID  string
	Qty      float64
	IsAsk    int
	Price    int64
}

func TestCancelOrder(t *testing.T) {
	api, err := makeAPI()
	if err != nil {
		t.Error(err)
	}

	tests := []CancelOrderTest{
		{
			Currency: BTC,
			OrderID:  "sdaslkdjslkajdskd",
			Qty:      0.005,
			IsAsk:    1,
			Price:    1000000,
		},
	}

	for _, v := range tests {
		cancelResp, err := api.CancelOrder(v.Currency, v.OrderID, v.Qty, v.IsAsk, v.Price)
		if err != nil {
			t.Error(err)
		}

		if cancelResp == nil {
			t.Error(err)
		}

		//log.Println(cancelResp)
	}
}

type TestOrder struct {
	Currency string
	Price    int64
	Qty      float64
}

func TestLimitBuy(t *testing.T) {
	api, err := makeAPI()
	if err != nil {
		t.Error(err)
	}

	tests := []TestOrder{
		{
			Currency: BTC,
			Price:    1000000,
			Qty:      0.005,
		},
	}

	for _, v := range tests {
		buyResp, err := api.LimitBuy(v.Currency, v.Price, v.Qty)
		if err != nil {
			t.Error(err)
		}

		if buyResp == nil {
			t.Error(err)
		}

		//log.Println(buyResp)
	}
}

func TestLimitSell(t *testing.T) {
	api, err := makeAPI()
	if err != nil {
		t.Error(err)
	}

	tests := []TestOrder{
		{
			Currency: BTC,
			Price:    1000000,
			Qty:      0.005,
		},
	}

	for _, v := range tests {
		sellResp, err := api.LimitSell(v.Currency, v.Price, v.Qty)
		if err != nil {
			t.Error(err)
		}

		if sellResp == nil {
			t.Error(err)
		}

		//log.Println(sellResp)
	}
}

func TestLimitOrders(t *testing.T) {
	api, err := makeAPI()
	if err != nil {
		t.Error(err)
	}

	tests := []test{
		{Coin: BTC},
	}

	for _, v := range tests {
		listResp, err := api.LimitOrders(v.Coin)
		if err != nil {
			t.Error(err)
		}

		if listResp == nil {
			t.Error(err)
		}

		//log.Println(listResp)
	}
}

func TestCompleteOrders(t *testing.T) {
	api, err := makeAPI()
	if err != nil {
		t.Error(err)
	}

	tests := []test{
		{Coin: BTC},
	}

	for _, v := range tests {
		listResp, err := api.CompleteOrders(v.Coin)
		if err != nil {
			t.Error(err)
		}

		if listResp == nil {
			t.Error(err)
		}

		//log.Println(listResp)
	}
}
