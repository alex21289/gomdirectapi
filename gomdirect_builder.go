package gomdirectapi

import "github.com/alex21289/merkur"

type ClientCredentials struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	AccessToken  string `json:"access_token"`
}

type sessionBuilder struct {
	credentials ClientCredentials
}

type SessionBuilder interface {
	Build() *comdirectSession
}

func NewBuilder(credentials ClientCredentials) SessionBuilder {
	builder := &sessionBuilder{
		credentials: credentials,
	}
	return builder
}

func (sb *sessionBuilder) Build() *comdirectSession {
	httpClient := merkur.NewBuilder().Build()
	session := comdirectSession{
		builder:    sb,
		httpClient: httpClient,
	}
	return &session
}
