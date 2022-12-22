package models

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type BalanceView struct {
	Account struct {
		AccountNumber  int         `json:"account_number"  db:"account_number"`
		Balance        float64     `json:"balance"  db:"balance"`
		Password       null.String `json:"password, omitempty"  db:"password" `
		CreatedAt      time.Time   `json:"created_at"  db:"created_at"`
		UpdatedAt      time.Time   `json:"updated_at"  db:"updated_at"`
		BalanceDetails []struct {
			TransactionsToken null.String `json:"transactions_token"  db:"transactions_token"`
			Value             null.Float  `json:"value"  db:"value"`
			Description       null.String `json:"description"  db:"description"`
			TypeTransactions  null.String `json:"type_transactions"  db:"type_transactions"`
			CreatedAt         null.Time   `json:"created_at"  db:"created_at"`
			UpdatedAt         null.Time   `json:"updated_at"  db:"updated_at"`
		} `json:"balance_details"`
	}
}

type Balance struct {
	TransactionsToken null.String `json:"transactions_token"  db:"transactions_token"`
	Value             null.Float  `json:"value"  db:"value"`
	Description       null.String `json:"description"  db:"description"`
	TypeTransactions  null.String `json:"type_transactions"  db:"type_transactions"`
	CreatedAt         null.Time   `json:"created_at"  db:"created_at"`
	UpdatedAt         null.Time   `json:"updated_at"  db:"updated_at"`
}
