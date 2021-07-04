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

func NewSession(credentials ClientCredentials) *ComdirectSession {
	httpClient := merkur.NewBuilder().Build()
	return &ComdirectSession{
		httpClient:  httpClient,
		Credentials: credentials,
		Session: auth.Session{
			ClientID:     credentials.ClientID,
			ClientSecret: credentials.ClientSecret,
		},
	}

}

func GetSessionFromJson(path string) (*ComdirectSession, error) {
	httpClient := merkur.NewBuilder().Build()
	var session auth.Session

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(file, &session)
	if err != nil {
		return nil, err
	}

	cs := ComdirectSession{
		httpClient: httpClient,
		Session:    session,
	}
	return &cs, nil
}
