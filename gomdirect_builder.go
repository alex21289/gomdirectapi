package gomdirectapi

import (
	"encoding/json"
	"io/ioutil"

	"github.com/alex21289/gomdirectapi/auth"
	"github.com/alex21289/merkur"
)

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
		AuthClient: auth.Client{
			ClientID:     sb.credentials.ClientID,
			ClientSecret: sb.credentials.ClientSecret,
		},
	}
	return &session
}

func GetSessionFormJson(path string) (*comdirectSession, error) {
	httpClient := merkur.NewBuilder().Build()
	var session auth.Client

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &session)
	if err != nil {
		return nil, err
	}

	cs := comdirectSession{
		httpClient: httpClient,
		AuthClient: session,
	}
	return &cs, nil
}
