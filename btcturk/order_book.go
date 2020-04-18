package btcturk

import "fmt"

type OrderBook struct {
	TimeStamp float64    `json:"timestamp"`
	Bids      [][]string `json:"bids"`
	Asks      [][]string `json:"asks"`
}

func (c *Client) OrderBook() (OrderBook, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/v2/orderbook?%s", c.params.Encode()), nil)

	var response OrderBook
	if _, err = c.do(req, &response); err != nil {
		return OrderBook{}, err
	}

	return response, nil
}
