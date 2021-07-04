package auth

import (
	"github.com/google/uuid"
)

type Session struct {
	ClientID         string    `json:"client_id"`
	ClientSecret     string    `json:"client_secret"`
	AccessToken      string    `json:"access_token"`
	TokenType        string    `json:"token_type"`
	RefreshToken     string    `json:"refresh_token"`
	ExpiresIn        int64     `json:"expires_in"`
	Scope            string    `json:"scope"`
	SessionUUID      string    `json:"session_uuid"`
	RequestID        string    `json:"request_id"`
	SessionID        uuid.UUID `json:"session_id"`
	QSession         string    `json:"qSession"`
	Tan              string    `json:"tan"`
	ChallengeID      string    `json:"challenge_id"`
	Bpid             int64     `json:"bpid"`
	Kdnr             string    `json:"kdnr"`
	KontaktID        int64     `json:"kontaktId"`
	Identifier       string    `json:"identifier"`
	SessionTanActive bool      `json:"sessionTanActive"`
	Activated2FA     bool      `json:"activated2FA"`
}

type RefreshSession struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

// Authentication holds the Response of Auth()
type Authentication struct {
	AccessToken  string `json:"access_token"`
	Bpid         int64  `json:"bpid"`
	ExpiresIn    int64  `json:"expires_in"`
	Kdnr         string `json:"kdnr"`
	KontaktID    int64  `json:"kontaktId"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
}

type SessionStatus struct {
	Identifier       string `json:"identifier"`
	SessionTanActive bool   `json:"sessionTanActive"`
	Activated2FA     bool   `json:"activated2FA"`
}

type AuthenticationInfo struct {
	ID             string   `json:"id"`
	Typ            string   `json:"typ"`
	AvailableTypes []string `json:"availableTypes"`
	Link           link     `json:"link"`
}

type link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
	Type   string `json:"type"`
}
