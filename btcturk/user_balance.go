package btcturk

type Balance struct {
	Asset     string `json:"asset"`
	AssetName string `json:"assetname"`
	Balance   string `json:"balance"`
	Locked    string `json:"locked"`
	Free      string `json:"free"`
}

func (c *Client) Balance() ([]Balance, error) {
	req, err := c.newRequest("GET", "/api/v1/users/balances", nil)
	if err != nil {
		return []Balance{}, err
	}
	if err := c.auth(req); err != nil {
		return []Balance{}, err
	}

	var response []Balance
	if _, err = c.do(req, &response); err != nil {
		return []Balance{}, err
	}

	return response, nil
}
