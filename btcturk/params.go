package btcturk

import (
	"fmt"
	"strconv"
)

// Limit Transactions
func (c *Client) Limit(v int) *Client { return c.addParamInt("limit", v) }

// Offset Transactions
func (c *Client) Offset(v int) *Client { return c.addParamInt("offset", v) }

// Sort Transactions
func (c *Client) Sort(v string) *Client {
	c.params.Add("sort", v)
	return c
}

// OrderID Cancel
func (c *Client) OrderID(v int) *Client { return c.addParamInt("id", v) }

// Quantity Buy or Sell
func (c *Client) Quantity(v float64) *Client { return c.addParamFloat("quantity", v) }

// Price Buy or Sell
func (c *Client) Price(v float64) *Client { return c.addParamFloat("price", v) }

// StopPrice Buy or Sell
func (c *Client) StopPrice(v float64) *Client { return c.addParamFloat("stopPrice", v) }

// NewOrderClientID Buy or Sell
func (c *Client) NewOrderClientID(v string) *Client {
	c.params.Add("newOrderClientId", v)
	return c
}

// OrderMethod Buy or Sell ( "limit", "market", "stoplimit" )
func (c *Client) OrderMethod(v string) *Client {
	c.params.Add("orderMethod", v)
	return c
}

// PairSymbol must be (Buy or Sell), Open Orders, Trades, Order Book
func (c *Client) PairSymbol(v string) *Client {
	c.params.Add("pairSymbol", v)
	return c
}

// Custom Param
func (c *Client) AddCustomParam(k string, v string) *Client {
	c.params.Add(k, v)
	return c
}

func (c *Client) addParamInt(key string, value int) *Client {
	c.params.Add(key, strconv.Itoa(value))
	return c
}

func (c *Client) addParamFloat(key string, value float64) *Client {
	c.params.Add(key, fmt.Sprintf("%f", value))
	return c
}
