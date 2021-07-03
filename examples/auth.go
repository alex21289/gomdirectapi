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
	viper.AddConfigPath("F:\\workspace\\github.com\\gomdirectapi\\examples")
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
		log.Println("Validierungsfehler")
		log.Println(err)
		os.Exit(1)
	}
	time.Sleep(time.Second * 10)
	// confirmSession()
	err = session.Activate()
	if err != nil {
		log.Fatal(err)
	}
	err = session.OAuth2()
	if err != nil {
		log.Println(err)
		os.Exit(21)
	}
	client, err := gomdirectapi.GetClient(session)
	if err != nil {
		log.Println(err)
		os.Exit(22)
	}

	accounts, err := client.GetAccounts()
	if err != nil {
		log.Println(err)
		os.Exit(23)
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
