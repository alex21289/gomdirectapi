package gomdirectapi

import "gomdirectapi/auth"

// NewClient creates a new Struct of Credentials
// Expects client_scret, client_Id, username and password
// returns Credentails struct
func NewClient(cID string, cS string, u string, p string) (c auth.Client) {
	c.ClientID = cID
	c.ClientSecret = cS
	c.Username = u
	c.Password = p
	return c
}
