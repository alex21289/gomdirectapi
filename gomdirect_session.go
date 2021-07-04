package gomdirectapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alex21289/gomdirectapi/auth"
	"github.com/alex21289/merkur"
	"github.com/alex21289/merkur/mmime"
	"github.com/google/uuid"
)

type comdirectSession struct {
	builder    *sessionBuilder
	httpClient merkur.Client

	Session        auth.Session
	RefreshSession auth.RefreshSession
	SessionStatus  auth.SessionStatus
	Authentication auth.Authentication
	AuthClient     auth.Client
}

type Session interface {
	Authenticate() error
	GetSession() error
	Validate() error
	Activate() error
	OAuth2() error
	Refresh() error
}

func (cs *comdirectSession) Authenticate() error {

	payload := "client_id=" + cs.builder.credentials.ClientID +
		"&client_secret=" + cs.builder.credentials.ClientSecret +
		"&grant_type=password&username=" + cs.builder.credentials.Username +
		"&password=" + cs.builder.credentials.Password

	headers := make(http.Header)
	headers.Set(mmime.HeaderContentType, mmime.ContentTypeXFormUrlencoded)
	headers.Set("Accept", mmime.ContentTypeJson)

	response, err := cs.httpClient.Post(TokenURL, payload, headers)
	if err != nil {
		return err
	}
	if response.StatusCode >= 400 {
		var httpError auth.Error
		response.UnmarshalJson(httpError)
		err := errors.New(httpError.ErrorDescription)
		return err
	}

	if err := response.UnmarshalJson(&cs.AuthClient); err != nil {
		return err
	}

	return nil
}

func (cs *comdirectSession) GetSession() error {
	now := time.Now().Format("20060201150405")
	cs.AuthClient.SessionID = uuid.New()
	cs.AuthClient.RequestID = now[len(now)-9 : len(now)]

	headers := make(http.Header)
	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", cs.AuthClient.SessionID, cs.AuthClient.RequestID)
	headers.Set("Accept", mmime.ContentTypeJson)
	headers.Set("Authorization", "Bearer "+cs.AuthClient.AccessToken)
	headers.Set("x-http-request-info", requestInfo)
	headers.Set(mmime.HeaderContentType, mmime.ContentTypeJson)

	response, err := cs.httpClient.Get(SessionURL, headers)
	if err != nil {
		return err
	}
	if response.StatusCode >= 400 {
		var sessionErr auth.Error
		response.UnmarshalJson(&sessionErr)
		err := errors.New(sessionErr.ErrorDescription)
		return err
	}

	qSession, ok := response.Cookies["qSession"]
	if !ok {
		err := errors.New("qSession cookie is missing")
		return err
	}
	cs.AuthClient.QSession = qSession.Value

	var sessions []auth.SessionStatus
	err = response.UnmarshalJson(&sessions)
	if err != nil {
		return err
	}
	if len(sessions) == 0 {
		return errors.New("creating session failed")
	}

	cs.AuthClient.SessionUUID = sessions[0].Identifier
	cs.AuthClient.Identifier = sessions[0].Identifier
	cs.SessionStatus.Identifier = sessions[0].Identifier
	cs.AuthClient.SessionTanActive = true
	cs.AuthClient.SessionTanActive = true
	cs.SessionStatus.Activated2FA = true
	cs.SessionStatus.SessionTanActive = true

	return nil
}

func (cs *comdirectSession) Validate() error {
	url := fmt.Sprintf(ValidateURL, cs.AuthClient.SessionUUID)

	headers := make(http.Header)
	accessToken := fmt.Sprintf("Bearer %s", cs.AuthClient.AccessToken)
	qSession := fmt.Sprintf("qSession=%s", cs.AuthClient.QSession)
	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", cs.AuthClient.SessionID, cs.AuthClient.RequestID)

	headers.Set("Accept", "application/json")
	headers.Set("Authorization", accessToken)
	headers.Set("x-http-request-info", requestInfo)
	headers.Set(mmime.HeaderContentType, mmime.ContentTypeJson)
	headers.Set("Cookie", qSession)

	response, err := cs.httpClient.Post(url, cs.SessionStatus, headers)
	if err != nil {
		return err
	}
	if response.StatusCode >= 400 {
		var err auth.ValidationError
		response.UnmarshalJson(&err)
		return errors.New(err.Messages[0]["message"])
	}

	var authInfo auth.AuthenticationInfo
	authInfoHeader := response.Headers.Get("x-once-authentication-info")
	if authInfoHeader == "" {
		return errors.New("missing x-once-authentication-info header")
	}

	err = json.Unmarshal([]byte(authInfoHeader), &authInfo)
	if err != nil {
		return err
	}
	cs.AuthClient.ChallengeID = authInfo.ID

	return nil
}

func (cs *comdirectSession) Activate() error {

	url := fmt.Sprintf(ActivateURL, cs.AuthClient.SessionUUID)

	accessToken := fmt.Sprintf("Bearer %s", cs.AuthClient.AccessToken)
	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", cs.AuthClient.SessionID, cs.AuthClient.RequestID)
	qSession := fmt.Sprintf("qSession=%s", cs.AuthClient.QSession)
	authInfo := fmt.Sprintf("{\"id\":\"%s\"}", cs.AuthClient.ChallengeID)
	headers := make(http.Header)

	headers.Set("Accept", "application/json")
	headers.Set("Authorization", accessToken)
	headers.Set("x-http-request-info", requestInfo)
	headers.Set("Content-Type", "application/json")
	headers.Set("x-once-authentication-info", authInfo)
	headers.Set("x-once-authentication", "")
	headers.Set("Cookie", qSession)

	// TODO: Handle response and error
	cs.httpClient.Patch(url, cs.SessionStatus, headers)
	log.Println("Succesfully Activate Session")

	return nil
}

func (cs *comdirectSession) OAuth2() error {

	payloadString := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=cd_secondary&token=%s", cs.builder.credentials.ClientID, cs.builder.credentials.ClientSecret, cs.AuthClient.AccessToken)
	headers := make(http.Header)
	qSession := fmt.Sprintf("qSession=%s", cs.AuthClient.QSession)
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Accept", "application/json")
	headers.Set("Cookie", qSession)

	result, err := cs.httpClient.Post(OAuth2URL, payloadString, headers)
	if err != nil {
		return err
	}
	if result.StatusCode >= 400 {
		return errors.New(result.String())
	}

	err = result.UnmarshalJson(&cs.Authentication)
	if err != nil {
		return err
	}
	cs.AuthClient.AccessToken = cs.Authentication.AccessToken
	cs.AuthClient.RefreshToken = cs.Authentication.RefreshToken

	// cs.AuthClient.ClientID = cs.builder.credentials.ClientID
	// cs.AuthClient.ClientSecret = cs.builder.credentials.ClientSecret

	return nil
}

func (cs *comdirectSession) Refresh() error {
	url := "https://api.comdirect.de/oauth/token"

	payloadString := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=refresh_token&refresh_token=%s", cs.AuthClient.ClientID, cs.AuthClient.ClientSecret, cs.AuthClient.RefreshToken)

	qSession := fmt.Sprintf("qSession=%s", cs.AuthClient.QSession)

	headers := make(http.Header)
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Accept", "application/json")
	headers.Set("Cookie", qSession)

	response, err := cs.httpClient.Post(url, payloadString, headers)
	if err != nil {
		return err
	}
	if response.StatusCode >= 400 {
		return errors.New(response.String() + " Status: " + response.Status)
	}
	err = response.UnmarshalJson(&cs.RefreshSession)
	if err != nil {
		return err
	}

	cs.AuthClient.AccessToken = cs.RefreshSession.AccessToken
	cs.AuthClient.RefreshToken = cs.RefreshSession.RefreshToken
	cs.AuthClient.ExpiresIn = time.Now().UTC().Add(time.Duration(cs.RefreshSession.ExpiresIn) * time.Second).Unix()

	return nil
}

func (cs *comdirectSession) Revoke() error {
	headers := make(http.Header)
	qSession := fmt.Sprintf("qSession=%s", cs.AuthClient.QSession)
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Accept", "application/json")
	headers.Set("Cookie", qSession)

	result, err := cs.httpClient.Delete(RevokeURL, headers)
	if err != nil {
		return err
	}
	if result.StatusCode >= 400 {
		return errors.New(result.String())
	}
	return nil
}

func (cs *comdirectSession) SaveToJson(path string) error {

	fileName := "session.json"
	filePath := filepath.Join(path, fileName)
	session, err := json.MarshalIndent(cs.AuthClient, "", "  ")
	if err != nil {
		return err
	}
	ioutil.WriteFile(filePath, session, 0644)
	return nil
}
