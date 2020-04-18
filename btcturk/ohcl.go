package btcturk

import "fmt"

type Ohcl struct {
	Pair                  string  `json:"pair"`
	TimeStamp             float64 `json:"time"`
	Open                  float64 `json:"open"`
	High                  float64 `json:"high"`
	Low                   float64 `json:"low"`
	Close                 float64 `json:"close"`
	Volume                float64 `json:"volume"`
	Average               float64 `json:"total"`
	DailyChangeAmount     float64 `json:"average"`
	DailyChangePercentage float64 `json:"dailyChangeAmount"`
}

func (c *Client) OhclData() ([]Ohcl, error) {
	//TODO: ohcl url
	req, err := c.newRequest("GET", fmt.Sprintf("/api/ohlcdata?%s", c.params.Encode()), nil)

	var response []Ohcl
	if _, err = c.do(req, &response); err != nil {
		return []Ohcl{}, err
	}

	return response, nil
}
