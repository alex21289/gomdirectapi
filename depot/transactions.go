package depot

type DepotTransactionApiResponse struct {
	Paging Paging         `json:"paging"`
	Values []Transactions `json:"values"`
}

type Staticdata struct {
	Notation               string `json:"notation"`
	Currency               string `json:"currency"`
	Instrumenttype         string `json:"instrumentType"`
	Priipsrelevant         bool   `json:"priipsRelevant"`
	Kidavailable           bool   `json:"kidAvailable"`
	Shippingwaiverrequired bool   `json:"shippingWaiverRequired"`
	Fundredemptionlimited  bool   `json:"fundRedemptionLimited"`
}

type Executionprice struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Transactionvalue struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
type Instrument struct {
	Instrumentid string     `json:"instrumentId"`
	Wkn          string     `json:"wkn"`
	Isin         string     `json:"isin"`
	Name         string     `json:"name"`
	Shortname    string     `json:"shortName"`
	Staticdata   Staticdata `json:"staticData"`
}
type Transactions struct {
	Transactionid        interface{}      `json:"transactionId"`
	Bookingstatus        string           `json:"bookingStatus"`
	Bookingdate          string           `json:"bookingDate"`
	Settlementdate       interface{}      `json:"settlementDate"`
	Businessdate         string           `json:"businessDate"`
	Quantity             Quantity         `json:"quantity"`
	Instrumentid         string           `json:"instrumentId"`
	Instrument           Instrument       `json:"instrument,omitempty"`
	Executionprice       Executionprice   `json:"executionPrice"`
	Transactionvalue     Transactionvalue `json:"transactionValue"`
	Transactiondirection string           `json:"transactionDirection"`
	Transactiontype      string           `json:"transactionType"`
	Fxrate               interface{}      `json:"fxRate"`
}
