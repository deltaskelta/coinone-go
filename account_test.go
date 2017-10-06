package coinone

import (
	"testing"
)

func TestGetBalance(t *testing.T) {
	api, err := makeAPI()
	if err != nil {
		t.Error(err)
	}

	balanceResp, err := api.GetBalance()
	if err != nil {
		t.Error(err)
	}

	if balanceResp == nil {
		t.Error(err)
	}

	//log.Println(balanceResp)
}
