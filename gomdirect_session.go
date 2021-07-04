package gomdirectapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alex21289/gomdirectapi/auth"
	"github.com/alex21289/merkur"
	"github.com/alex21289/merkur/formdata"
	"github.com/alex21289/merkur/mmime"
	"github.com/google/uuid"
)

type ComdirectSession struct {
	Credentials ClientCredentials
	httpClient  merkur.Client

	Session        auth.Session
	RefreshSession auth.RefreshSession
	SessionStatus  auth.SessionStatus
	Authentication auth.Authentication
}

type Session interface {
	Authenticate() error
	GetSession() error
	Validate() error
	Activate() error
	OAuth2() error
	Refresh() error
}

func (cs *ComdirectSession) Authenticate() error {

	payload := formdata.NewFormData()
	payload.Set("client_id", cs.Credentials.ClientID)
	payload.Set("client_secret", cs.Credentials.ClientSecret)
	payload.Set("grant_type", "password")
	payload.Set("username", cs.Credentials.Username)
	payload.Set("password", cs.Credentials.Password)

	headers := make(http.Header)
	headers.Set(mmime.HeaderContentType, mmime.ContentTypeXFormUrlencoded)
	headers.Set("Accept", mmime.ContentTypeJson)

	response, err := cs.httpClient.Post(TokenURL, payload, headers)
	if err := handleErr(*response, err); err != nil {
		return err
	}

	if err := response.UnmarshalJson(&cs.Session); err != nil {
		return err
	}

	return nil
}

func (cs *ComdirectSession) GetSession() error {
	now := time.Now().Format("20060201150405")
	cs.Session.SessionID = uuid.New()
	cs.Session.RequestID = now[len(now)-9 : len(now)]

	headers := make(http.Header)
	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", cs.Session.SessionID, cs.Session.RequestID)
	headers.Set("Accept", mmime.ContentTypeJson)
	headers.Set("Authorization", "Bearer "+cs.Session.AccessToken)
	headers.Set("x-http-request-info", requestInfo)
	headers.Set(mmime.HeaderContentType, mmime.ContentTypeJson)

	response, err := cs.httpClient.Get(SessionURL, headers)
	if err := handleErr(*response, err); err != nil {
		return err
	}

	qSession, ok := response.Cookies["qSession"]
	if !ok {
		err := errors.New("qSession cookie is missing")
		return err
	}
	cs.Session.QSession = qSession.Value

	var sessions []auth.SessionStatus
	err = response.UnmarshalJson(&sessions)
	if err != nil {
		return err
	}
	if len(sessions) == 0 {
		return errors.New("creating session failed")
	}

	cs.Session.SessionUUID = sessions[0].Identifier
	cs.Session.Identifier = sessions[0].Identifier
	cs.Session.SessionTanActive = true

	cs.SessionStatus.Identifier = sessions[0].Identifier
	cs.SessionStatus.SessionTanActive = true
	cs.SessionStatus.Activated2FA = true

	return nil
}

func (cs *ComdirectSession) Validate() error {
	url := fmt.Sprintf(ValidateURL, cs.Session.SessionUUID)

	headers := make(http.Header)
	accessToken := fmt.Sprintf("Bearer %s", cs.Session.AccessToken)
	qSession := fmt.Sprintf("qSession=%s", cs.Session.QSession)
	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", cs.Session.SessionID, cs.Session.RequestID)

	headers.Set("Accept", "application/json")
	headers.Set("Authorization", accessToken)
	headers.Set("x-http-request-info", requestInfo)
	headers.Set(mmime.HeaderContentType, mmime.ContentTypeJson)
	headers.Set("Cookie", qSession)

	response, err := cs.httpClient.Post(url, cs.SessionStatus, headers)
	if err := handleErr(*response, err); err != nil {
		return err
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
	cs.Session.ChallengeID = authInfo.ID

	return nil
}

func (cs *ComdirectSession) Activate() error {

	url := fmt.Sprintf(ActivateURL, cs.Session.SessionUUID)

	accessToken := fmt.Sprintf("Bearer %s", cs.Session.AccessToken)
	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", cs.Session.SessionID, cs.Session.RequestID)
	qSession := fmt.Sprintf("qSession=%s", cs.Session.QSession)
	authInfo := fmt.Sprintf("{\"id\":\"%s\"}", cs.Session.ChallengeID)
	headers := make(http.Header)

	headers.Set("Accept", "application/json")
	headers.Set("Authorization", accessToken)
	headers.Set("x-http-request-info", requestInfo)
	headers.Set("Content-Type", "application/json")
	headers.Set("x-once-authentication-info", authInfo)
	headers.Set("x-once-authentication", "")
	headers.Set("Cookie", qSession)

	response, err := cs.httpClient.Patch(url, cs.SessionStatus, headers)
	if err := handleErr(*response, err); err != nil {
		return err
	}

	return nil
}

func (cs *ComdirectSession) OAuth2() error {

	payload := formdata.NewFormData()
	payload.Set("client_id", cs.Credentials.ClientID)
	payload.Set("client_secret", cs.Credentials.ClientSecret)
	payload.Set("grant_type", "cd_secondary")
	payload.Set("token", cs.Session.AccessToken)

	headers := make(http.Header)
	qSession := fmt.Sprintf("qSession=%s", cs.Session.QSession)
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Accept", "application/json")
	headers.Set("Cookie", qSession)

	response, err := cs.httpClient.Post(OAuth2URL, payload, headers)
	if err := handleErr(*response, err); err != nil {
		return err
	}

	err = response.UnmarshalJson(&cs.Authentication)
	if err != nil {
		return err
	}
	cs.Session.AccessToken = cs.Authentication.AccessToken
	cs.Session.RefreshToken = cs.Authentication.RefreshToken

	return nil
}

func (cs *ComdirectSession) Refresh() error {
	url := "https://api.comdirect.de/oauth/token"

	payload := formdata.NewFormData()
	payload.Set("client_id", cs.Credentials.ClientID)
	payload.Set("client_secret", cs.Credentials.ClientSecret)
	payload.Set("grant_type", "refresh_token")
	payload.Set("refresh_token", cs.Session.RefreshToken)

	qSession := fmt.Sprintf("qSession=%s", cs.Session.QSession)

	headers := make(http.Header)
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Accept", "application/json")
	headers.Set("Cookie", qSession)

	response, err := cs.httpClient.Post(url, payload, headers)
	if err := handleErr(*response, err); err != nil {
		return err
	}

	err = response.UnmarshalJson(&cs.RefreshSession)
	if err != nil {
		return err
	}

	cs.Session.AccessToken = cs.RefreshSession.AccessToken
	cs.Session.RefreshToken = cs.RefreshSession.RefreshToken
	cs.Session.ExpiresIn = time.Now().UTC().Add(time.Duration(cs.RefreshSession.ExpiresIn) * time.Second).Unix()

	return nil
}

func (cs *ComdirectSession) Revoke() error {
	headers := make(http.Header)
	qSession := fmt.Sprintf("qSession=%s", cs.Session.QSession)
	headers.Set("Content-Type", "application/x-www-form-urlencoded")
	headers.Set("Accept", "application/json")
	headers.Set("Cookie", qSession)

	response, err := cs.httpClient.Delete(RevokeURL, headers)
	if err := handleErr(*response, err); err != nil {
		return err
	}

	return nil
}

func (cs *ComdirectSession) SaveToJson(path string) error {

	fileName := "session.json"
	filePath := filepath.Join(path, fileName)
	session, err := json.MarshalIndent(cs.Session, "", "  ")
	if err != nil {
		return err
	}
	ioutil.WriteFile(filePath, session, 0644)
	return nil
}
