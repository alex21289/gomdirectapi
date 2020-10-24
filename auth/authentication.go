package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

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

// Client provides the login credentials
type Client struct {
	ClientID     string
	ClientSecret string
	Username     string
	Password     string
	AccessToken  string
	TokenType    string
	RefreshToken string
	Scope        string
}

// Error is uses to parse an error Response
type Error struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// NewClient creates a new Struct of Credentials
// Expects client_scret, client_Id, username and password
// returns Credentails struct
func NewClient(cID string, cS string, u string, p string) (c Client) {
	c.ClientID = cID
	c.ClientSecret = cS
	c.Username = u
	c.Password = p
	return c
}

// Auth authenticates with the Credentials to the comdirect API
// and sets the tokens to the client struct
// Returns the response as Authentication struct
func (c *Client) Auth() (auth Authentication, err error) {
	// Comdirect API URL (2020-10-24)
	url := "https://api.comdirect.de/oauth/token"

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
	}

	// Add Headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	// Send the request to the given url
	res, err := client.Do(req)
	if err != nil {
		log.Println("The request fails")
		log.Println(err)
	}
	defer res.Body.Close()

	// Read response body
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Reading body went wrong...")
		log.Println(err)
	}

	// Convert body in byteSlice (rb = response body)
	rb := []byte(string(body))

	fmt.Println(body)
	// Parse json response into Authentication Struct
	jerr := json.Unmarshal(rb, &auth)
	// jerr, cant use err ?!
	if jerr != nil {
		log.Println("JSON could not decoded")
		log.Println(jerr)
	}

	fmt.Println("Error:", jerr)
	c.AccessToken = auth.AccessToken
	c.RefreshToken = auth.RefreshToken
	c.TokenType = auth.TokenType
	c.Scope = auth.Scope

	// Returns Authentication Struct
	return auth, jerr
}
