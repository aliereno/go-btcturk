package btcturk

import (
	"fmt"
)

type UserTransactions struct {
	Price             float64 `json:"price"`
	NumeratorSymbol   string  `json:"numeratorSymbol"`
	DenominatorSymbol string  `json:"denominatorSymbol"`
	OrderType         string  `json:"orderType"`
	ID                string  `json:"id"`
	Timestamp         float64 `json:"timestamp"`
	Amount            float64 `json:"amount"`
	Fee               float64 `json:"fee"`
	Tax               float64 `json:"tax"`
}

func (c *Client) UserTransactions() ([]UserTransactions, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/v1/users/transactions/trade?%s", c.params.Encode()), nil)
	if err != nil {
		return []UserTransactions{}, err
	}
	if err := c.auth(req); err != nil {
		return []UserTransactions{}, err
	}

	var response []UserTransactions
	if _, err = c.do(req, &response); err != nil {
		return []UserTransactions{}, err
	}

	return response, nil
}
