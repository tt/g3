package api

import (
	"errors"
)

type Account struct {
	ID string `json:"id"`
}

var accountTable accounts

type accounts []Account

func (accountTable *accounts) Insert(account Account) {
	*accountTable = accounts(append(*accountTable, account))
}

func (accountTable accounts) Find(id string) (*Account, error) {
	for _, account := range accountTable {
		if account.ID == id {
			return &account, nil
		}
	}

	return nil, errors.New("account not found")
}
