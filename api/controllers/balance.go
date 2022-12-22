package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/oklog/ulid"
	"github.com/orenhapeba1/estudy-api-golang-bank/databases"
	"github.com/orenhapeba1/estudy-api-golang-bank/models"
	"github.com/orenhapeba1/estudy-api-golang-bank/services"
	"log"
	"math/rand"
	"time"
)

func genUlid() string {
	t := time.Now().UTC()
	entropy := rand.New(rand.NewSource(t.UnixNano()))
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return id.String()
}

func BalanceView(c *gin.Context) {
	code := 200
	msgreturn := ""
	AccountReceive := c.Param("account") // RECEBE ID PELA URL

	token := services.TokenBearer(c.GetHeader("Authorization"))

	if token == "" {
		c.JSON(401, gin.H{
			"error": "Token Invalid",
		})
		return
	} else {
		Db := databases.GetDatabase()
		BalanceDetailsAccountViewList := []models.BalanceView{}
		//BalanceViewList := []models.Balance{}
		//AccountBalanceList := []models.AccountBalance{}

		qry := ""
		varConsulta := ""

		if AccountReceive != "" {
			qry = "SELECT a.account_id, a.account_number, ab.balance, a.created_at, a.updated_at FROM accounts a LEFT JOIN account_balance ab ON ab.account_id = a.account_id WHERE a.account_number = ? LIMIT 1"
			varConsulta = AccountReceive
		} else {
			qry = "SELECT a.account_id, a.account_number, ab.balance, a.created_at, a.updated_at FROM accounts a LEFT JOIN account_balance ab ON ab.account_id = a.account_id WHERE a.token = ? LIMIT 1"
			varConsulta = token
		}

		rows, err := Db.Queryx(qry, varConsulta)

		if err != nil {
			log.Fatal(err)
			code = 500
			msgreturn = "Erro MySQL:" + err.Error()
		}
		for rows.Next() {
			BalanceDetailsAccountView := models.BalanceView{}
			AccountBalance := models.AccountBalance{}

			err = rows.StructScan(&AccountBalance)

			AccountId := AccountBalance.AccountId
			BalanceDetailsAccountView.Account.AccountNumber = AccountBalance.AccountNumber
			BalanceDetailsAccountView.Account.Balance = AccountBalance.Balance
			BalanceDetailsAccountView.Account.Password = AccountBalance.Password
			BalanceDetailsAccountView.Account.CreatedAt = AccountBalance.CreatedAt
			BalanceDetailsAccountView.Account.UpdatedAt = AccountBalance.UpdatedAt

			/*leitura extrato*/

			qry = "SELECT t.transactions_token,t.value,t.description,t.type_transactions,t.created_at,t.updated_at FROM transactions t WHERE t.account_id = ? "
			rows1, err1 := Db.Queryx(qry, AccountId)

			if err1 != nil {
				log.Fatal(err1)
				code = 500
				msgreturn = "Erro MySQL:" + err1.Error()
			}

			BalanceDetails := models.Balance{}
			for rows1.Next() {
				err = rows1.StructScan(&BalanceDetails)
				BalanceDetailsAccountView.Account.BalanceDetails = append(BalanceDetailsAccountView.Account.BalanceDetails, BalanceDetails) //ADICIONA AO ARRAY

			}
			/*leitura extrato*/

			BalanceDetailsAccountViewList = append(BalanceDetailsAccountViewList, BalanceDetailsAccountView) //ADICIONA AO ARRAY

		}

		if err != nil {
			log.Fatal(err)
			code = 500
			msgreturn = "Erro Leitura Linhas MySQL:" + err.Error()
		}

		defer rows.Close()

		if code != 200 {
			c.JSON(code, gin.H{
				"error": msgreturn,
			})
			return
		} else {
			c.JSON(200, BalanceDetailsAccountViewList)
		}

	}

}

func BalanceInsert(c *gin.Context) {
	code := 200
	msgreturn := ""
	qry := ""
	varConsulta := ""

	Db := databases.GetDatabase()

	AccountReceive := c.Param("account") // RECEBE ID PELA URL
	token := services.TokenBearer(c.GetHeader("Authorization"))

	var p models.Balance
	err := c.ShouldBindJSON(&p)

	if err != nil {
		code = 400
		msgreturn = "cannot bind JSON: " + err.Error()

	}

	if AccountReceive != "" {
		qry = "SELECT a.account_id, a.account_number, ab.balance, a.created_at, a.updated_at FROM accounts a LEFT JOIN account_balance ab ON ab.account_id = a.account_id WHERE a.account_number = ? LIMIT 1"
		varConsulta = AccountReceive
	} else {
		qry = "SELECT a.account_id, a.account_number, ab.balance, a.created_at, a.updated_at FROM accounts a LEFT JOIN account_balance ab ON ab.account_id = a.account_id WHERE a.token = ? LIMIT 1"
		varConsulta = token
	}

	rows, err := Db.Queryx(qry, varConsulta)

	if err != nil {
		log.Fatal(err)
		code = 500
		msgreturn = "Erro MySQL:" + err.Error()
	}

	AccountBalance := models.AccountBalance{}
	for rows.Next() {
		err = rows.StructScan(&AccountBalance)

	}

	if err != nil {
		log.Fatal(err)
		code = 500
		msgreturn = "Erro StructScan:" + err.Error()
	}

	TypeTransactions := ""
	qryupdate := ""
	fmt.Println(p.TypeTransactions.String)
	if p.TypeTransactions.String == "e" {
		TypeTransactions = "RECEIVED"
		qryupdate = "UPDATE account_balance SET balance=balance + ? WHERE account_id = ?"
	} else if p.TypeTransactions.String == "s" {
		TypeTransactions = "PAID_OUT"
		qryupdate = "UPDATE account_balance SET balance=balance - ? WHERE account_id = ?"
	} else {
		TypeTransactions = "AWAITING_RISK_ANALYSIS"
	}
	_, err1 := Db.Queryx(qryupdate, p.Value, AccountBalance.AccountId)

	if err1 != nil {
		log.Fatal(err1)
		code = 500
		msgreturn = "Erro MySQL:" + err1.Error()
	}

	qry = "INSERT INTO transactions (transactions_token,account_id,value,description,type_transactions) VALUES (?,?,?,?,?)"
	stmt, err := Db.Prepare(qry)
	_, err = stmt.Exec(genUlid(), AccountBalance.AccountId, p.Value, p.Description, TypeTransactions)

	if err != nil {
		log.Fatal(err)
		code = 500
		msgreturn = "Erro MySQL:" + err.Error()
	}

	if code != 200 {
		c.JSON(code, gin.H{
			"error": msgreturn,
		})
		return
	} else {
		c.JSON(200, "completed transaction")
	}

}
