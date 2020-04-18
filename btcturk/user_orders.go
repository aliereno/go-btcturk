package btcturk

import (
	"fmt"
	"strings"
)

const (
	BTCTRY string = "BTCTRY"
	ETHTRY string = "ETHTRY"

	MarketOrder     int = 1
	LimitOrder      int = 0
	StopMarketOrder int = 3
)

type OpenOrders struct {
	ID         string  `json:"id"`
	DateTime   string  `json:"datetime"`
	Type       string  `json:"type"`
	Price      float64 `json:"price"`
	Amount     float64 `json:"amount"`
	PairSymbol string  `json:"PairSymbol"`
}

type OrderType struct {
	ID       string  `json:"id"`
	DateTime string  `json:"datetime"`
	Type     string  `json:"type"`
	Price    float64 `json:"price"`
	Amount   float64 `json:"amount"`
}

func (c *Client) OpenOrders() ([]OpenOrders, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/openOrders?%s", c.params.Encode()), nil)
	if err := c.auth(req); err != nil {
		return []OpenOrders{}, err
	}

	var response []OpenOrders
	if _, err = c.do(req, &response); err != nil {
		return []OpenOrders{}, err
	}

	return response, nil
}

func (c *Client) Buy() (OrderType, error) {
	c.params.Add("OrderType", "0")
	c.body = strings.NewReader(c.params.Encode())

	req, err := c.newRequest("POST", "/api/exchange", c.body)
	if err := c.auth(req); err != nil {
		return OrderType{}, err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

	var response OrderType
	if _, err = c.do(req, &response); err != nil {
		return OrderType{}, err
	}

	return response, nil
}

func (c *Client) Sell() (OrderType, error) {
	c.params.Add("OrderType", "1")
	c.body = strings.NewReader(c.params.Encode())

	req, err := c.newRequest("POST", "/api/exchange", c.body)
	if err := c.auth(req); err != nil {
		return OrderType{}, err
	}

	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")

	var response OrderType
	if _, err = c.do(req, &response); err != nil {
		return OrderType{}, err
	}

	return response, nil
}

func (c *Client) CancelOrder() (bool, error) {
	c.body = strings.NewReader(c.params.Encode())

	req, err := c.newRequest("POST", "/api/cancelOrder", c.body)
	if err := c.auth(req); err != nil {
		return false, err
	}

	var response struct {
		Result bool `json:"result,omitempty"`
	}
	if _, err = c.do(req, &response); err != nil {
		return false, err
	}

	return response.Result, nil
}
