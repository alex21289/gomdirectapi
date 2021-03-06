package report

type ReportApiResponse struct {
	Paging     Paging     `json:"paging"`
	Aggregated Aggregated `json:"aggregated"`
	Values     []Values   `json:"values"`
}
type Paging struct {
	Index   int `json:"index"`
	Matches int `json:"matches"`
}
type Balanceeur struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Availablecashamounteur struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Aggregated struct {
	Balanceeur             Balanceeur             `json:"balanceEUR"`
	Availablecashamounteur Availablecashamounteur `json:"availableCashAmountEUR"`
}
type Accounttype struct {
	Key  string `json:"key"`
	Text string `json:"text"`
}
type Creditlimit struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Account struct {
	Accountid        string      `json:"accountId"`
	Accountdisplayid string      `json:"accountDisplayId"`
	Currency         string      `json:"currency"`
	Accounttype      Accounttype `json:"accountType"`
	Iban             string      `json:"iban"`
	Creditlimit      Creditlimit `json:"creditLimit"`
}
type Balance struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Availablecashamount struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type AccountBalance struct {
	Account                Account                `json:"account"`
	Accountid              string                 `json:"accountId"`
	Balance                Balance                `json:"balance"`
	Balanceeur             Balanceeur             `json:"balanceEUR"`
	Availablecashamount    Availablecashamount    `json:"availableCashAmount"`
	Availablecashamounteur Availablecashamounteur `json:"availableCashAmountEUR"`
	Depotid                string                 `json:"depotId"`
	Depot                  Depot                  `json:"depot"`
	Datelastupdate         string                 `json:"dateLastUpdate"`
	Prevdayvalue           Prevdayvalue           `json:"prevDayValue"`
}
type Depot struct {
	Depotid                    string        `json:"depotId"`
	Depotdisplayid             string        `json:"depotDisplayId"`
	Clientid                   string        `json:"clientId"`
	Defaultsettlementaccountid string        `json:"defaultSettlementAccountId"`
	Settlementaccountids       []interface{} `json:"settlementAccountIds"`
}
type Prevdayvalue struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}

type Values struct {
	Productid            string         `json:"productId"`
	Producttype          string         `json:"productType"`
	Targetclientid       string         `json:"targetClientId"`
	Clientconnectiontype string         `json:"clientConnectionType"`
	Balance              AccountBalance `json:"balance,omitempty"`
}
