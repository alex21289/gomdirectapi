package gomdirectapi

const (
	TokenURL   string = "https://api.comdirect.de/oauth/token"
	SessionURL string = "https://api.comdirect.de/api/session/clients/user/v1/sessions"
	// ValidateURL
	// Must use with fmt.Sprintf to pass the sessionUUID
	ValidateURL string = "https://api.comdirect.de/api/session/clients/user/v1/sessions/%s/validate"
	ActivateURL string = "https://api.comdirect.de/api/session/clients/user/v1/sessions/%s"
	OAuth2URL   string = "https://api.comdirect.de/oauth/token"
	RevokeURL   string = "https://api.comdirect.de/oauth/revoke"
)
