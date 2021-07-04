package auth

import (
	"github.com/google/uuid"
)

type Session struct {
	ClientID     string    `json:"client_id"`
	ClientSecret string    `json:"client_secret"`
	AccessToken  string    `json:"access_token"`
	RequestID    string    `json:"request:id"`
	SessionID    uuid.UUID `json:"session_id"`
	QSession     string    `json:"qSession"`
	RefreshToken string    `json:"refresh_token"`
	ExpiresIn    int64     `json:"expires_in"`
	Expires      int64     `json:"expires"`
}

type RefreshSession struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	Scope        string `json:"scope"`
}

// Client provides the login credentials
type Client struct {
	ClientID         string    `json:"client_id"`
	ClientSecret     string    `json:"client_secret"`
	Username         string    `json:"username"`
	Password         string    `json:"password"`
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

// Error is uses to parse an error Response
type Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type ValidationError struct {
	Code     string              `json:"code"`
	Messages []map[string]string `json:"messages"`
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
