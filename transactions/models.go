package transactions

// TODO: replace type interface{}

type TransactionAPIResponse struct {
	Paging struct {
		Index   int `json:"index"`
		Matches int `json:"matches"`
	} `json:"paging"`
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

type Transaction struct {
	Reference     string `json:"reference"`
	BookingStatus string `json:"bookingStatus"`
	BookingDate   string `json:"bookingDate"`
	Amount        struct {
		Value string `json:"value"`
		Unit  string `json:"unit"`
	} `json:"amount"`
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
