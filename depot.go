package gomdirectapi

import (
	"fmt"

	"github.com/alex21289/gomdirectapi/depot"
)

// GetDepot returns the information of all depots
func (c *Client) GetDepot() (*depot.DepotApiResponse, error) {
	response, err := c.http.Get(DepotURL)
	if err := handleErr(*response, err); err != nil {
		return nil, err
	}

	var d depot.DepotApiResponse
	if err = response.UnmarshalJson(&d); err != nil {
		return nil, err
	}

	return &d, nil
}

// GetDepotProtfolio returns the portfolio of the given depotID
func (c *Client) GetDepotProtfolio(depotID string) (*depot.DepotPortfolioApiResponse, error) {
	url := fmt.Sprintf(DepotPortfolioURL, depotID)
	response, err := c.http.Get(url)
	if err := handleErr(*response, err); err != nil {
		return nil, err
	}

	var dp depot.DepotPortfolioApiResponse
	if err = response.UnmarshalJson(&dp); err != nil {
		return nil, err
	}

	return &dp, nil
}

// GetDepotPosition returns detailed information of the given position
func (c *Client) GetDepotPosition(depotID string, positionID string) (*depot.Position, error) {
	url := fmt.Sprintf(DepotPositionURL, depotID, positionID)
	response, err := c.http.Get(url)
	if err := handleErr(*response, err); err != nil {
		return nil, err
	}

	var dp depot.Position
	if err = response.UnmarshalJson(&dp); err != nil {
		return nil, err
	}

	return &dp, nil
}

// GetDepotTransactions returns all transactions of the given depotID
func (c *Client) GetDepotTransactions(depotID string) (*depot.DepotTransactionApiResponse, error) {
	url := fmt.Sprintf(DepotTransactionsURL, depotID)
	response, err := c.http.Get(url)
	if err := handleErr(*response, err); err != nil {
		return nil, err
	}

	var dt depot.DepotTransactionApiResponse
	if err = response.UnmarshalJson(&dt); err != nil {
		return nil, err
	}

	return &dt, nil
}
