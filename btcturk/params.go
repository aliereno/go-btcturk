package btcturk

import (
	"fmt"
	"strconv"
)

// Limit Transactions
func (c *Client) Limit(v int) *Client { return c.addParamInt("limit", v) }

// Offset Transactions
func (c *Client) Offset(v int) *Client { return c.addParamInt("offest", v) }

// OrderID Cancel
func (c *Client) OrderID(v int) *Client { return c.addParamInt("id", v) }

// OrderMethod Buy or Sell
func (c *Client) OrderMethod(v int) *Client { return c.addParamInt("OrderMethod", v) }

// PricePrecision Buy or Sell
func (c *Client) PricePrecision(v int) *Client { return c.addParamPrecision("PricePrecision", v) }

// Amount Buy or Sell
func (c *Client) Amount(v int) *Client { return c.addParamInt("Amount", v) }

// AmountPrecision Buy or Sell
func (c *Client) AmountPrecision(v int) *Client { return c.addParamPrecision("AmountPrecision", v) }

// Total Buy or Sell
func (c *Client) Total(v int) *Client { return c.addParamInt("Total", v) }

// TotalPrecision Buy or Sell
func (c *Client) TotalPrecision(v int) *Client { return c.addParamPrecision("TotalPrecision", v) }

// TriggerPrice Buy or Sell
func (c *Client) TriggerPrice(v int) *Client { return c.addParamInt("TriggerPrice", v) }

// Price Buy or Sell
func (c *Client) Price(v int) *Client { return c.addParamInt("Price", v) }

// Last Trades, OHCL Data Params (Not required)
func (c *Client) Last(v int) *Client { return c.addParamInt("last", v) }

// PairSymbol must be (Buy or Sell), Open Orders, Trades, Order Book
func (c *Client) PairSymbol(v string) *Client {
	c.params.Add("PairSymbol", v)
	return c
}

// TriggerPricePrecision Buy or Sell
func (c *Client) TriggerPricePrecision(v int) *Client {
	return c.addParamPrecision("TriggerPricePrecision", v)
}

// Sort Transactions
func (c *Client) Sort(v string) *Client {
	c.params.Add("sort", v)
	return c
}

func (c *Client) addParamInt(key string, value int) *Client {
	c.params.Add(key, strconv.Itoa(value))
	return c
}

func (c *Client) addParamPrecision(key string, value int) *Client {
	c.params.Add(key, fmt.Sprintf("%03d", value))
	return c
}
