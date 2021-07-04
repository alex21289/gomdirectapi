package gomdirectapi

const (
	BaseURL    string = "https://api.comdirect.de/"
	TokenURL   string = BaseURL + "oauth/token"
	SessionURL string = BaseURL + "api/session/clients/user/v1/sessions"

	// ValidateURL use with fmt.Sprintf to pass the sessionUUID
	ValidateURL string = BaseURL + "api/session/clients/user/v1/sessions/%s/validate"

	// ActivateURL use with fmt.Sprintf to pass the sessionUUID
	ActivateURL string = BaseURL + "api/session/clients/user/v1/sessions/%s"
	OAuth2URL   string = BaseURL + "oauth/token"
	RevokeURL   string = BaseURL + "oauth/revoke"
)

// Depot
const (
	DepotURL string = BaseURL + "/brokerage/clients/user/v3/depots"
)
