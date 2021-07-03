# Go Comdirect API

This will be a simple Go wrapper for the private comdirect API.

The package will provide

- Authentication
- Read transactions and balances

## Example Authentication

### main.go

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alex21289/gomdirectapi"
	"github.com/alex21289/gomdirectapi/account"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("<path-to-credentials>")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	creds := gomdirectapi.ClientCredentials{
		ClientID:     viper.GetString("client_id"),
		ClientSecret: viper.GetString("client_secret"),
		Username:     viper.GetString("username"),
		Password:     viper.GetString("Password"),
		AccessToken:  viper.GetString("access_token"),
	}

	// Create a new session instance
	session := gomdirectapi.NewBuilder(creds).Build()

	// Authentication flow step by step
	if err := session.Authenticate(); err != nil {
		log.Fatal(err)
	}
	if err := session.GetSession(); err != nil {
		log.Fatal(err)
	}
	err := session.Validate()
	if err != nil {
		log.Fatal(err)
	}
	// Confirm session with mobile device
	confirmSession()

	if err := session.Activate(); err != nil {
		log.Fatal(err)
	}
if 	err := session.OAuth2(); err != nil {
	log.Fatal(err)
}

// Get Client from session
client, err := gomdirectapi.GetClient(session)
	if err != nil {
		log.Fatal(err)
	}


// Create an instance to query accounts
	accounts, err := account.GetAccounts(client)
	if err != nil {
		log.Println(err)
		os.Exit(23)
	}
	log.Println(accounts.Values)

}

func confirmSession() {
	fmt.Print("Press Enter after you confirmed the Session on your Mobile Device... ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
}

```
