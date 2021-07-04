package gomdirectapi

const (
	BaseURL    string = "https://api.comdirect.de"
	TokenURL   string = BaseURL + "/oauth/token"
	SessionURL string = BaseURL + "/api/session/clients/user/v1/sessions"

	// ValidateURL use with fmt.Sprintf to pass the sessionUUID
	ValidateURL string = BaseURL + "/api/session/clients/user/v1/sessions/%s/validate"

	// ActivateURL use with fmt.Sprintf to pass the sessionUUID
	ActivateURL string = BaseURL + "/api/session/clients/user/v1/sessions/%s"
	OAuth2URL   string = BaseURL + "/oauth/token"
	RevokeURL   string = BaseURL + "/oauth/revoke"

	// Depot
	DepotURL             string = BaseURL + "/api/brokerage/clients/user/v3/depots"
	DepotPortfolioURL    string = BaseURL + "/api/brokerage/v3/depots/%s/positions"
	DepotPositionURL     string = BaseURL + "/api/brokerage/v3/depots/%s/positions/%s"
	DepotTransactionsURL string = BaseURL + "/api/brokerage/v3/depots/%s/transactions"

	// Postbox
	PostboxURL         string = BaseURL + "/api/messages/clients/user/v2/documents"
	PostboxDocumentURL string = BaseURL + "/api/messages/v2/documents/%s"

	// Report
	ReportURL string = BaseURL + "/api/reports/participants/user/v1/allbalances"
)
