package btcturk

// https://docs.btcturk.com/#ticker
type Ticker struct {
	Pair              string  `json:"pair"`
	PairNormalized    string  `json:"pairNormalized"`
	Timestamp         int64   `json:"timestamp"`
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

// If pairSymbol is not set, ticker for all pairs will be returned in a json array.
// Or
// GET ?pairSymbol=BTC_TRY
// Or
// GET ?symbol=USDT
func (c *Client) Ticker() ([]Ticker, error) {
	req, err := c.newRequest("GET", "/api/v2/ticker", nil)
	if err != nil {
		return []Ticker{}, err
	}
	var response []Ticker
	if _, err = c.do(req, &response); err != nil {
		return []Ticker{}, err
	}
	return response, nil
}
