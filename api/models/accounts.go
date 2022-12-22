package models

import (
	"gopkg.in/guregu/null.v4"
	"time"
)

type Accounts_Token struct {
	Token string `json:"token"  db:"token"`
}

type Accounts_Password_Validation struct {
	Password string `json:"password"  db:"password"`
}

type Account struct {
	AccountId     null.Int    `json:"account_id"  db:"account_id"`
	AccountNumber int         `json:"account_number"  db:"account_number"`
	Password      null.String `json:"password"  db:"password"`
	Token         string      `json:"token"  db:"token"`
	CreatedAt     time.Time   `json:"created_at"  db:"created_at"`
	UpdatedAt     time.Time   `json:"updated_at"  db:"updated_at"`
}
