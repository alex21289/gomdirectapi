package gomdirectapi

import (
	"errors"
	"fmt"

	"github.com/alex21289/gomdirectapi/account"
	"github.com/alex21289/gomdirectapi/transactions"
)

func (c *Client) GetAccounts() (*account.AccountAPIResponse, error) {
	AccountURL := "https://api.comdirect.de/api/banking/clients/user/v2/accounts/balances"
	response, err := c.http.Get(AccountURL)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, errors.New(response.String() + "Status " + response.Status)
	}

	var accounts account.AccountAPIResponse
	response.UnmarshalJson(&accounts)
	return &accounts, nil
}

func (c *Client) GetAccountByUUID(uuid string) (*account.Account, error) {
	var SingleAccountURL = "https://api.comdirect.de/api/banking/v2/accounts/%s/balances"
	url := fmt.Sprintf(SingleAccountURL, uuid)
	response, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, errors.New(response.String() + "Status " + response.Status)
	}
	var account account.Account
	if err := response.UnmarshalJson(&account); err != nil {
		return nil, err
	}
	return &account, nil
}

func (c *Client) GetTransactions(accountUUID string) (*transactions.TransactionAPIResponse, error) {
	var TransactionURL = "https://api.comdirect.de/api/banking/v1/accounts/%s/transactions"

	url := fmt.Sprintf(TransactionURL, accountUUID)

	response, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode >= 400 {
		return nil, errors.New(response.String() + " Status: " + response.Status)
	}

	var transactions transactions.TransactionAPIResponse
	err = response.UnmarshalJson(&transactions)
	if err != nil {
		return nil, err
	}
	return &transactions, nil
}
