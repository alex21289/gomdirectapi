package auth

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

const (
	TokenURL   = "https://api.comdirect.de/oauth/token"
	SessionURL = "https://api.comdirect.de/api/session/clients/user/v1/sessions"
	// ValidateURL
	// Must use with fmt.Sprintf to pass the sessionUUID
	ValidateURL = "https://api.comdirect.de/api/session/clients/user/v1/sessions/%s/validate"
	ActivateURL = "https://api.comdirect.de/api/session/clients/user/v1/sessions/%s"
	OAuth2URL   = "https://api.comdirect.de/oauth/token"
)

// NewClient creates a new Struct of Credentials
// Expects client_scret, client_Id, username and password
// returns Credentails struct
func NewClient(clientID string, clientSecret string, username string, password string) (c Client) {
	c.ClientID = clientID
	c.ClientSecret = clientSecret
	c.Username = username
	c.Password = password
	return c
}

func (c *Client) NewSession() (*Session, error) {
	err := c.Auth()
	if err != nil {
		return nil, err
	}
	c.GetSession()
	c.Validate()
	confirmSession()
	c.Activate()
	c.OAuth2()

	var session = Session{
		ClientID:     c.ClientID,
		ClientSecret: c.ClientSecret,
		AccessToken:  c.AccessToken,
		RequestID:    c.RequestID,
		SessionID:    c.SessionID,
		QSession:     c.QSession,
		RefreshToken: c.RefreshToken,
		ExpiresIn:    c.ExpiresIn,
		Expires:      time.Now().UTC().Add(time.Duration(c.ExpiresIn) * time.Second).Unix(),
	}

	return &session, nil
}

// Auth authenticates with the Credentials to the comdirect API
// and sets the tokens to the client struct
// Returns the response as Authentication struct
func (c *Client) Auth() error {
	var respError Error
	// Comdirect API URL (2020-10-24)
	url := TokenURL

	method := "POST"

	// Creates the payload with Credential Struct
	payload := strings.NewReader("client_id=" + c.ClientID +
		"&client_secret=" + c.ClientSecret +
		"&grant_type=password&username=" + c.Username +
		"&password=" + c.Password)

	// Initialize a http Client
	client := &http.Client{}

	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		log.Println("Creating request fails...")
		log.Println(err)
		return err
	}

	// Add Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	// Send the request to the given url
	res, err := client.Do(req)
	if err != nil {
		log.Println("The request fails")
		log.Println(err)
		return err
	}
	defer res.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Reading body went wrong...")
		log.Println(err)
		return err
	}

	// Parse json response into Authentication Struct
	jerr := json.Unmarshal(body, &c)
	if jerr != nil {
		log.Println("JSON could not decoded")
		log.Println(jerr)
		return err
	}

	// Parse json error response into Error Struct
	respErr := json.Unmarshal(body, &respError)
	if respErr != nil {
		log.Println("JSON could not decoded")
		log.Println(respErr)
		return err
	}

	if (Error{} != respError) {
		err = errors.New(respError.ErrorDescription)
		log.Println("Authentication Error")
		return err
	}
	log.Println("Succesfully get AuthToken")

	return nil
}

func (c *Client) GetSession() {

	var sessions []SessionStatus

	url := SessionURL
	method := "GET"
	now := time.Now().Format("20060201150405")
	c.SessionID = uuid.New()
	c.RequestID = now[len(now)-9 : len(now)]

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	accessToken := fmt.Sprintf("Bearer %s", c.AccessToken)
	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", c.SessionID, c.RequestID)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", accessToken)
	req.Header.Add("x-http-request-info", requestInfo)
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// c.QSession = getSessionCookie("qSession", res)

	for _, cookie := range res.Cookies() {
		if cookie.Name == "qSession" {
			log.Println("Get Cookie:", cookie.Value)
			c.QSession = cookie.Value
		}
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = json.Unmarshal(body, &sessions)
	if err != nil {
		log.Panic("Error while Unmarshal the session")
	}
	session := &sessions[0]
	c.SessionUUID = session.Identifier
	log.Println("Succesfully create Session")
}

func (c *Client) Validate() {

	url := fmt.Sprintf(ValidateURL, c.SessionUUID)
	method := "POST"

	// Create new Session
	session := SessionStatus{
		Identifier:       c.SessionUUID,
		SessionTanActive: true,
		Activated2FA:     true,
	}

	// Marshal Session Struct to JSON
	jsonPayload, err := json.Marshal(session)
	if err != nil {
		log.Panic("Error while marshal Session")
	}

	// Decode JSON bytes to Reader
	payload := strings.NewReader(string(jsonPayload))
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	accessToken := fmt.Sprintf("Bearer %s", c.AccessToken)
	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", c.SessionID, c.RequestID)
	qSession := fmt.Sprintf("qSession=%s", c.QSession)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", accessToken)
	req.Header.Add("x-http-request-info", requestInfo)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Cookie", qSession)

	// Send Request
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	// Get Response Header
	var authInfo AuthenticationInfo
	authInfoHeader := res.Header.Get("x-once-authentication-info")

	err = json.Unmarshal([]byte(authInfoHeader), &authInfo)
	if err != nil {
		log.Panic("Error while unmarshal AuthInfo Header!")
	}

	c.ChallengeID = authInfo.ID
	if c.ChallengeID == "" {
		log.Panic("Got no Challenge ID")
	}

	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Println("Succesfully validate Session")
}

func (c *Client) Activate() {
	url := fmt.Sprintf(ActivateURL, c.SessionUUID)
	method := "PATCH"

	// Create new Session
	session := SessionStatus{
		Identifier:       c.SessionUUID,
		SessionTanActive: true,
		Activated2FA:     true,
	}

	// Marshal Session Struct to JSON
	jsonPayload, err := json.Marshal(session)
	if err != nil {
		log.Panic("Error while marshal Session")
	}

	// Decode JSON bytes to Reader
	payload := strings.NewReader(string(jsonPayload))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	accessToken := fmt.Sprintf("Bearer %s", c.AccessToken)
	requestInfo := fmt.Sprintf("{\"clientRequestId\":{\"sessionId\":\"%s\",\"requestId\":\"%s\"}}", c.SessionID, c.RequestID)
	qSession := fmt.Sprintf("qSession=%s", c.QSession)
	authInfo := fmt.Sprintf("{\"id\":\"%s\"}", c.ChallengeID)

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Authorization", accessToken)
	req.Header.Add("x-http-request-info", requestInfo)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("x-once-authentication-info", authInfo)
	req.Header.Add("x-once-authentication", "")
	req.Header.Add("Cookie", qSession)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()
	log.Println("Succesfully Activate Session")
}

func (c *Client) OAuth2() {
	url := OAuth2URL
	method := "POST"

	payloadString := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=cd_secondary&token=%s", c.ClientID, c.ClientSecret, c.AccessToken)
	payload := strings.NewReader(payloadString)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return
	}

	qSession := fmt.Sprintf("qSession=%s", c.QSession)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cookie", qSession)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Convert body in byteSlice (rb = response body)
	rb := []byte(string(body))

	// Parse json response into Authentication Struct
	jerr := json.Unmarshal(rb, &c)
	if jerr != nil {
		log.Println("JSON could not decoded")
		log.Println(jerr)
	}
	fmt.Println("Authentication sucess")
}

func confirmSession() {
	fmt.Print("Press Enter after confirm the Session on your Mobile Device")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
}

func (s *Session) Refresh() error {

	var refreshSession RefreshSession
	url := "https://api.comdirect.de/oauth/token"
	method := "POST"

	payloadString := fmt.Sprintf("client_id=%s&client_secret=%s&grant_type=refresh_token&refresh_token=%s", s.ClientID, s.ClientSecret, s.RefreshToken)
	payload := strings.NewReader(payloadString)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		fmt.Println(err)
		return err
	}

	qSession := fmt.Sprintf("qSession=%s", s.QSession)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Cookie", qSession)

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err
	}

	json.Unmarshal(body, &refreshSession)

	s.AccessToken = refreshSession.AccessToken
	s.RefreshToken = refreshSession.RefreshToken
	s.Expires = time.Now().UTC().Add(time.Duration(refreshSession.ExpiresIn) * time.Second).Unix()
	for _, cookie := range res.Cookies() {
		if cookie.Name == "qSession" {
			log.Println("Get Cookie:", cookie.Value)
			s.QSession = cookie.Value
		}
	}

	return nil
}
