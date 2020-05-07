package btcturk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
)

const (
	BTCTRY   string = "BTCTRY"
	ETHBTC   string = "ETHBTC"
	ETHTRY   string = "ETHTRY"
	XRPTRY   string = "XRPTRY"
	LTCTRY   string = "LTCTRY"
	USDTTRY  string = "USDTTRY"
	BTCUSDT  string = "BTCUSDT"
	ETHUSDT  string = "ETHUSDT"
	XRPUSDT  string = "XRPUSDT"
	LTCUSDT  string = "LTCUSDT"
	XLMTRY   string = "XLMTRY"
	LTCBTC   string = "LTCBTC"
	XRPBTC   string = "XRPBTC"
	XLMBTC   string = "XLMBTC"
	XLMUSDT  string = "XLMUSDT"
	NEOTRY   string = "NEOTRY"
	NEOBTC   string = "NEOBTC"
	NEOUSDT  string = "NEOUSDT"
	EOSTRY   string = "EOSTRY"
	EOSUSDT  string = "EOSUSDT"
	EOSBTC   string = "EOSBTC"
	DASHTRY  string = "DASHTRY"
	DASHBTC  string = "DASHBTC"
	DASHUSDT string = "DASHUSDT"
	LINKTRY  string = "LINKTRY"
	LINKUSDT string = "LINKUSDT"
	LINKBTC  string = "LINKBTC"
	ATOMTRY  string = "ATOMTRY"
	ATOMUSDT string = "ATOMUSDT"
	ATOMBTC  string = "ATOMBTC"
	XTZTRY   string = "XTZTRY"
	TRY      string = "TRY"
	BTC      string = "BTC"
	ETH      string = "ETH"
	XRP      string = "XRP"
	LTC      string = "LTC"
	USDT     string = "USDT"
	XLM      string = "XLM"
	NEO      string = "NEO"
	EOS      string = "EOS"
	DASH     string = "DASH"
	LINK     string = "LINK"
	ATOM     string = "ATOM"
	XTZ      string = "XTZ"

	MarketOrder    string = "market"
	LimitOrder     string = "limit"
	StopLimitOrder string = "stoplimit"
)

type OpenOrderModel struct {
	ID            int32   `json:"id"`
	Price         string  `json:"price"`
	Amount        string  `json:"amount"`
	Quantity      string  `json:"quantity"`
	PairSymbol    string  `json:"pairsymbol"`
	Type          string  `json:"type"`
	Method        string  `json:"method"`
	OrderClientID string  `json:"orderClientId"`
	DateTime      float64 `json:"datetime"`
	UpdateTime    float64 `json:"updateTime"`
	Status        string  `json:"status"`
}

type OpenOrderResult struct {
	Asks []OpenOrderModel `json:"asks"`
	Bids []OpenOrderModel `json:"bids"`
}

type OrderResult struct {
	ID                   int32  `json:"id"`
	DateTime             int32  `json:"datetime"`
	Type                 string `json:"type"`
	Method               string `json:"method"`
	Price                string `json:"price"`
	StopPrice            string `json:"stopPrice"`
	Quantity             string `json:"quantity"`
	PairSymbol           string `json:"pairSymbol"`
	PairSymbolNormalized string `json:"pairSymbolNormalized"`
	NewOrderClientID     string `json:"newOrderClientId"`
}

type OrderInput struct {
	Quantity         float64 `json:"quantity"`
	Price            float64 `json:"price"`
	StopPrice        float64 `json:"stopPrice"`
	NewOrderClientId string  `json:"newOrderClientId"`
	OrderMethod      string  `json:"orderMethod"`
	OrderType        string  `json:"orderType"`
	PairSymbol       string  `json:"pairSymbol"`
}

func (c *Client) OpenOrders() (OpenOrderResult, error) {
	jsonString, err := json.Marshal(c.params)
	if err != nil {
		return OpenOrderResult{}, err
	}
	req, err := c.newRequest("GET", "/api/v1/openOrders", bytes.NewBuffer(jsonString))
	if err != nil {
		return OpenOrderResult{}, err
	}
	if err := c.auth(req); err != nil {
		return OpenOrderResult{}, err
	}

	var response OpenOrderResult
	if _, err = c.do(req, &response); err != nil {
		return OpenOrderResult{}, err
	}

	return response, nil
}

func (c *Client) Buy(o *OrderInput) (OrderResult, error) {
	o.OrderType = "buy"
	jsonString, err := json.Marshal(o)
	if err != nil {
		return OrderResult{}, err
	}

	req, err := c.newRequest("POST", "/api/v1/order", bytes.NewBuffer(jsonString))
	if err != nil {
		return OrderResult{}, err
	}
	if err := c.auth(req); err != nil {
		return OrderResult{}, err
	}

	var response OrderResult
	if _, err = c.do(req, &response); err != nil {
		return OrderResult{}, err
	}

	return response, nil
}

func (c *Client) Sell(o *OrderInput) (OrderResult, error) {
	o.OrderType = "sell"
	jsonString, err := json.Marshal(o)
	if err != nil {
		return OrderResult{}, err
	}

	req, err := c.newRequest("POST", "/api/v1/order", bytes.NewBuffer(jsonString))
	if err != nil {
		return OrderResult{}, err
	}
	if err := c.auth(req); err != nil {
		return OrderResult{}, err
	}

	var response OrderResult
	if _, err = c.do(req, &response); err != nil {
		return OrderResult{}, err
	}

	return response, nil
}

func (c *Client) CancelOrder() (bool, error) {
	req, err := c.newRequest("DELETE", fmt.Sprintf("/api/v1/order?%s", c.params.Encode()), c.body)
	if err != nil {
		return false, err
	}
	if err := c.auth(req); err != nil {
		return false, err
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return false, err
	}

	defer func() {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
		c.clearRequest()
	}()

	var response = &GeneralResponse{}

	if json.NewDecoder(resp.Body).Decode(response) != nil {
		return false, err
	}

	if response.Success == true {
		return true, nil
	} else {
		return false, errors.New(*response.Message)
	}

}
