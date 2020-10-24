package account

// Accounts holds all accounts
type Accounts struct {
	accounts map[string]Account
}

// Account holds the Account information
type Account struct {
	ID     string
	Name   string
	amount float32
}


func NewAccount(id struct) ac Account {
	return ac
}