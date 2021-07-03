package account

type AccountAPIResponse struct {
	Paging Paging    `json:"paging"`
	Values []Account `json:"values"`
}

type Paging struct {
	Index   int `json:"index"`
	Matches int `json:"matches"`
}

type Account struct {
	AccountDetail          AccountDetail          `json:"account"`
	AccountID              string                 `json:"accountId"`
	Balance                Balance                `json:"balance"`
	BalanceEUR             BalanceEUR             `json:"balanceEUR"`
	AvailableCashAmount    AvailableCashAmount    `json:"availableCashAmount"`
	AvailableCashAmountEUR AvailableCashAmountEUR `json:"availableCashAmountEUR"`
}

type AccountDetail struct {
	AccountID        string      `json:"accountId"`
	AccountDisplayID string      `json:"accountDisplayId"`
	Currency         string      `json:"currency"`
	ClientID         string      `json:"clientId"`
	AccountType      AccountType `json:"accountType"`
	IBAN             string      `json:"iban"`
	CreditLimit      CreditLimit `json:"creditLimit"`
}

type Balance struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type BalanceEUR struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type AvailableCashAmount struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type AvailableCashAmountEUR struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type AccountType struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}

type CreditLimit struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
