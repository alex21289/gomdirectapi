package transactions

type TransactionAPIResponse struct {
	Paging     Paging `json:"paging"`
	Aggregated struct {
		Account                      interface{} `json:"account"`
		AccountID                    string      `json:"accountId"`
		BookingDateLatestTransaction string      `json:"bookingDateLatestTransaction"`
		ReferenceLatestTransaction   string      `json:"referenceLatestTransaction"`
		LatestTransactionIncluded    bool        `json:"latestTransactionIncluded"`
		PagingTimestamp              string      `json:"pagingTimestamp"`
	} `json:"aggregated"`
	Values []Transaction `json:"values"`
}

type Paging struct {
	Index   int `json:"index"`
	Matches int `json:"matches"`
}

type Transaction struct {
	Reference             string      `json:"reference"`
	BookingStatus         string      `json:"bookingStatus"`
	BookingDate           string      `json:"bookingDate"`
	Amount                Amount      `json:"amount"`
	Remitter              Holder      `json:"remitter"`
	Deptor                Holder      `json:"deptor"`
	Creditor              Holder      `json:"creditor"`
	ValutaDate            string      `json:"valutaDate"`
	DirectDebitCreditorID interface{} `json:"directDebitCreditorId"`
	DirectDebitMandateID  interface{} `json:"directDebitMandateId"`
	EndToEndReference     string      `json:"endToEndReference"`
	NewTransaction        bool        `json:"newTransaction"`
	RemittanceInfo        string      `json:"remittanceInfo"`
	TransactionType       struct {
		Key  string `json:"key"`
		Text string `json:"text"`
	} `json:"transactionType"`
}

type Holder struct {
	HolderName string `json:"holderName"`
	Iban       string `json:"iban"`
	Bic        string `json:"bic"`
}

type Bookings struct {
	RemittanceInfo string
	Amount         string
	Debiting       string
	Count          int
	Values         []Booking
}

type Booking struct {
	Reference         string `json:"reference"`
	EndToEndReference string `json:"endToEndReference"`

	// RemittanceInfo: Verwendungszweck
	RemittanceInfo string `json:"remittanceInfo"`
	BookingStatus  string `json:"bookingStatus"`
	BookingDate    string `json:"bookingDate"`
	Amount         Amount `json:"amount"`

	// Remitter: Buchungen
	Remitter Holder `json:"remitter"`

	// Deptor
	Deptor Holder `json:"deptor"`

	// Creditor: Überweisungen
	Creditor Holder `json:"creditor"`
}

type Amount struct {
	Value string `json:"value"`
	Unit  string `json:"unit"`
}
