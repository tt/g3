package api

type Account struct {
	ID string `json:"id"`
}

var accountTable accounts

type accounts []Account

func (accountTable *accounts) Insert(account Account) {
	*accountTable = accounts(append(*accountTable, account))
}
