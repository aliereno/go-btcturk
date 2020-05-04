package btcturk

import "fmt"

// https://docs.btcturk.com/#trades
type Trade struct {
	Pair           string  `json:"pair"`
	PairNormalized string  `json:"pairNormalized"`
	Numerator      string  `json:"numerator"`
	Denominator    string  `json:"denominator"`
	TimeStamp      float64 `json:"date"`
	TID            string  `json:"tid"`
	Price          string  `json:"price"`
	Amount         string  `json:"amount"`
	Side           string  `json:"side"`
}

// GET ?pairSymbol=BTC_TRY
// or
// GET ?pairSymbol=BTC_TRY&last=COUNT (Max. value for count parameter is 50)
func (c *Client) Trades() ([]Trade, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/v2/trades?%s", c.params.Encode()), nil)
	if err != nil {
		return []Trade{}, err
	}
	var response []Trade
	if _, err = c.do(req, &response); err != nil {
		return []Trade{}, err
	}

	return response, nil
}
