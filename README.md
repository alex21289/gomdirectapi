# Go Comdirect API

This will be a simple Go wrapper for the private comdirect API.

The package will provide

- Authentication
- Read transactions and balances

## Example Authentication
 I suggested to store your credentials in an .env file

### main.go

```go
package main

import (
	"fmt"
	"github.com/alex21289/gomdirectapi"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Try to load .env files")
	// loads values from .env into the system
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env files found")
	}
}

func main() {

	clientID, isID := os.LookupEnv("clientID")
	clientSecret, isSecret := os.LookupEnv("clientSecret")
	username, isUser := os.LookupEnv("username")
	password, isPass := os.LookupEnv("password")

	// Check wheather all credentials are given
	if !isID || !isSecret || !isUser || !isPass {
		log.Fatal("Missing credentials")
	}

	client := auth.NewClient(clientID, clientSecret, username, password)
	auth, err := client.Auth()

	if err != nil {
		log.Println("Authentication went wrong")
		log.Println(err)
	}

	// You can access the Access token from both
	fmt.Println("Access Token from client:", client.AccessToken)
	fmt.Println("Access Token from auth:", auth.AccessToken)
}
```
