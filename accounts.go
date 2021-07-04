package gomdirectapi

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alex21289/gomdirectapi/account"
	"github.com/alex21289/gomdirectapi/transactions"
	"github.com/alex21289/merkur"
)

// GetAccounts returns all accounts
func (c *Client) GetAccounts() (*account.AccountAPIResponse, error) {
	AccountURL := "https://api.comdirect.de/api/banking/clients/user/v2/accounts/balances"
	response, err := c.http.Get(AccountURL)
	if err := handleErr(*response, err); err != nil {
		return nil, err
	}

	var accounts account.AccountAPIResponse
	response.UnmarshalJson(&accounts)
	return &accounts, nil
}

// GetAccountByID returns the details of the given Account
func (c *Client) GetAccountByID(accountID string) (*account.Account, error) {
	var SingleAccountURL = "https://api.comdirect.de/api/banking/v2/accounts/%s/balances"
	url := fmt.Sprintf(SingleAccountURL, accountID)
	response, err := c.http.Get(url)
	if err := handleErr(*response, err); err != nil {
		return nil, err
	}

	var account account.Account
	if err := response.UnmarshalJson(&account); err != nil {
		return nil, err
	}
	return &account, nil
}

// GetTransactions returns the transactions of the given AccountID
func (c *Client) GetTransactions(accountID string) (*transactions.TransactionAPIResponse, error) {
	var TransactionURL = "https://api.comdirect.de/api/banking/v1/accounts/%s/transactions"

	url := fmt.Sprintf(TransactionURL, accountID)

	// First call to get paging-count
	response, err := c.http.Get(url)
	if err := handleErr(*response, err); err != nil {
		return nil, err
	}

	var transactions transactions.TransactionAPIResponse
	err = response.UnmarshalJson(&transactions)
	if err != nil {
		return nil, err
	}

	// Get all transactions
	params := merkur.NewParams()
	params.Add("paging-count", strconv.Itoa(transactions.Paging.Matches))
	response, err = c.http.GetQuery(url, params)
	if err := handleErr(*response, err); err != nil {
		return nil, err
	}

	err = response.UnmarshalJson(&transactions)
	if err != nil {
		return nil, err
	}

	return &transactions, nil
}

// GetBookings returns a struct with a cleaner slice of a cusotm transactions bookings struct.
func (c *Client) GetBookings(accountID string) (*transactions.Bookings, error) {
	ts, err := c.GetTransactions(accountID)
	if err != nil {
		return nil, err
	}

	var b []transactions.Booking
	var amount float64 = 0
	var debiting float64 = 0
	for _, t := range ts.Values {

		booking := transactions.Booking{
			Reference:         t.Reference,
			EndToEndReference: t.EndToEndReference,
			RemittanceInfo:    cleanRemittanceInfo(t.RemittanceInfo),
			BookingStatus:     t.BookingStatus,
			BookingDate:       t.BookingDate,
			Amount:            t.Amount,
			Remitter:          t.Remitter,
			Deptor:            t.Deptor,
			Creditor:          t.Creditor,
		}
		b = append(b, booking)
		val, err := strconv.ParseFloat(t.Amount.Value, 64)
		if err != nil {
			return nil, err
		}
		if val < 0 {
			debiting = debiting + val
		} else {
			amount = amount + val
		}
	}

	bookings := transactions.Bookings{
		RemittanceInfo: "",
		Values:         b,
		Count:          len(b),
		Amount:         fmt.Sprintf("%.2f", amount),
		Debiting:       fmt.Sprintf("%.2f", debiting),
	}

	return &bookings, nil
}

func cleanRemittanceInfo(remittanceInfo string) string {
	remittanceInfo = remittanceInfo[2:]
	remittanceInfo = strings.ReplaceAll(remittanceInfo, "02End-to-End-Ref.:                   03nicht angegeben", "")
	remittanceInfo = strings.ReplaceAll(remittanceInfo, "\t", "")
	remittanceInfo = strings.TrimSpace(remittanceInfo)

	return remittanceInfo
}
