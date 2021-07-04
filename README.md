# Go Comdirect API (WIP)

This is a simple Go wrapper for the private comdirect API.

The package will provide

- Authentication
- Refreshing session
- Read account transactions and balances
- Read depot and transactions
- Read reports
- Read and download documents from postbox

## Example Authentication

```go
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alex21289/gomdirectapi"
	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("<path to credential.json>")
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

	session := gomdirectapi.NewBuilder(creds).Build()

	// Get Session
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

	// Confirm the session with mobile device
	confirmSession()
	err = session.Activate()
	if err != nil {
		log.Fatal(err)
	}
	err = session.OAuth2()
	if err != nil {
		log.Fatal(err)
	}
	client, err := gomdirectapi.GetClient(session)
	if err != nil {
		log.Println(err)
	}

	accounts, err := client.GetAccounts()
	if err != nil {
		log.Println(err)
	}

	for _, a := range accounts.Values {
		log.Println("============")
		log.Println(a.AccountDetail.AccountType.Text)
		log.Println(a.Balance.Value + " €")
	}
}

func confirmSession() {
	fmt.Print("Press Enter after confirm the Session on your Mobile Device")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
}
```
