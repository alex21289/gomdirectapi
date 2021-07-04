package main

import (
	"log"
	"os"

	"github.com/alex21289/gomdirectapi"
)

func main() {

	session, err := gomdirectapi.GetSessionFromJson("..\\session.json")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	client, err := gomdirectapi.GetClient(session)
	if err != nil {
		log.Println(err)
		os.Exit(2)
	}

	err = session.Refresh()
	if err != nil {
		log.Fatal(err)
	}

	client, err = gomdirectapi.GetClient(session)
	if err != nil {
		log.Fatal(err)
	}

	err = session.SaveToJson("..")
	if err != nil {
		log.Fatal(err)
	}

	account, err := client.GetAccounts()
	if err != nil {
		log.Fatal(err)
	}
	log.Println(account.Values[0].AccountDetail.AccountType.Text)
	log.Println(account.Values[0].Balance)
	t, err := client.GetTransactions(account.Values[0].AccountDetail.AccountID)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(t.Values[0].EndToEndReference)
	log.Println(t.Values[0].Amount.Value, t.Values[0].Amount.Unit)
}
