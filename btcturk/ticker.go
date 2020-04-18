package btcturk

type Ticker struct {
	Pair              string  `json:"pair"`
	PairNormalized    string  `json:"pairNormalized"`
	Timestamp         float32 `json:"timestamp"`
	Last              float64 `json:"last"`
	High              float64 `json:"high"`
	Low               float64 `json:"low"`
	Bid               float64 `json:"bid"`
	Ask               float64 `json:"ask"`
	Open              float64 `json:"open"`
	Volume            float64 `json:"volume"`
	Average           float64 `json:"average"`
	Daily             float64 `json:"daily"`
	DailyPercent      float64 `json:"dailyPercent"`
	DenominatorSymbol string  `json:"denominatorSymbol"`
	NumeratorSymbol   string  `json:"numeratorSymbol"`
}

func (c *Client) Ticker() ([]Ticker, error) {
	req, err := c.newRequest("GET", "/api/v2/ticker", nil)
	var response []Ticker
	if _, err = c.do(req, &response); err != nil {
		return []Ticker{}, err
	}
	return response, nil
}
