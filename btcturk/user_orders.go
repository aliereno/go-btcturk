package btcturk

import (
	"bytes"
	"encoding/json"
	"fmt"
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

type OpenOrders struct {
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

type OrderType struct {
	ID       string  `json:"id"`
	DateTime string  `json:"datetime"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	Amount   float64 `json:"amount"`
}

func (c *Client) OpenOrders() ([]OpenOrders, error) {
	jsonString, err := json.Marshal(c.params)
	req, err := c.newRequest("GET", "/api/v1/openOrders", bytes.NewBuffer(jsonString))
	if err != nil {
		return []OpenOrders{}, err
	}
	if err := c.auth(req); err != nil {
		return []OpenOrders{}, err
	}

	var response []OpenOrders
	if _, err = c.do(req, &response); err != nil {
		return []OpenOrders{}, err
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

	var response GeneralResponse

	// TODO
	// API returns `"code":""`
	// my code expects `"code":0` an integer
	// so it will return error
	if _, err = c.do(req, &response); err != nil {
		return false, err
	}

	return response.Success, nil
}

func (c *Client) Buy() (OrderType, error) {
	c.params.Add("orderType", "buy")
	jsonString, err := json.Marshal(c.params)
	if err != nil {
		return OrderType{}, err
	}

	req, err := c.newRequest("POST", "/api/v1/order", bytes.NewBuffer(jsonString))
	if err != nil {
		return OrderType{}, err
	}
	if err := c.auth(req); err != nil {
		return OrderType{}, err
	}

	var response OrderType
	if _, err = c.do(req, &response); err != nil {
		return OrderType{}, err
	}

	return response, nil
}

func (c *Client) Sell() (OrderType, error) {
	c.params.Add("orderType", "sell")
	jsonString, err := json.Marshal(c.params)
	if err != nil {
		return OrderType{}, err
	}

	req, err := c.newRequest("POST", "/api/v1/order", bytes.NewBuffer(jsonString))
	if err != nil {
		return OrderType{}, err
	}
	if err := c.auth(req); err != nil {
		return OrderType{}, err
	}

	var response OrderType
	if _, err = c.do(req, &response); err != nil {
		return OrderType{}, err
	}

	return response, nil
}
