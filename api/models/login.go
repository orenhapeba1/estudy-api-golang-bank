package models

type Login struct {
	Account  int    `json:"account" db:"account_number"`
	Password string `json:"password" db:"password"`
}
