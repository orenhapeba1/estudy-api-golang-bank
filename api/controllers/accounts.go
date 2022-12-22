package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/orenhapeba1/estudy-api-golang-bank/databases"
	"github.com/orenhapeba1/estudy-api-golang-bank/models"
	"github.com/orenhapeba1/estudy-api-golang-bank/services"
	"log"
	"strconv"
)

func AccountView(c *gin.Context) {
	code := 200
	msgreturn := ""

	token := services.TokenBearer(c.GetHeader("Authorization"))

	if token == "" {
		c.JSON(401, gin.H{
			"error": "Token Invalid",
		})
		return
	} else {
		Db := databases.GetDatabase()
		AccountView := models.AccountBalance{}

		qry := "SELECT a.account_number, ab.balance, a.created_at, a.updated_at FROM accounts a LEFT JOIN account_balance ab ON ab.account_id = a.account_id WHERE a.token = ? LIMIT 1"
		rows, err := Db.Queryx(qry, token)

		if err != nil {
			log.Fatal(err)
			code = 500
			msgreturn = "Erro MySQL:" + err.Error()
		}
		for rows.Next() {
			err = rows.StructScan(&AccountView)
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
			c.JSON(200, AccountView)
		}

	}

}

func AccountInsert(c *gin.Context) {
	code := 200
	msgreturn := ""

	Db := databases.GetDatabase()
	AccountMax := models.AccountMax{}
	AccountView := models.AccountBalance{}

	var p models.AccountCreate
	err := c.ShouldBindJSON(&p)
	p.Password = services.SHA256Encoder(p.Password)

	qry := "SELECT MAX(account_number)+1 AS account_number FROM accounts"
	rows, err := Db.Queryx(qry)

	if err != nil {
		log.Fatal(err)
		code = 500
		msgreturn = "Erro MySQL:" + err.Error()
	}
	for rows.Next() {
		err = rows.StructScan(&AccountMax)
	}
	if err != nil {
		log.Fatal(err)
		code = 500
		msgreturn = "Erro Leitura Linhas MySQL:" + err.Error()
	}

	account_number := AccountMax.AccountNumber
	qryinser := "INSERT INTO accounts (account_number, password) VALUES (?,?)"

	stmt, err := Db.Prepare(qryinser)
	res, err := stmt.Exec(account_number, p.Password)

	if err != nil {
		code = 500
		msgreturn = "Erro Leitura Linhas MySQL:" + err.Error()
	}

	LastIDEmpresa, err := res.LastInsertId()

	qry2 := "SELECT a.account_number, ab.balance, a.created_at, a.updated_at FROM accounts a LEFT JOIN account_balance ab ON ab.account_id = a.account_id WHERE a.account_id = ? LIMIT 1"
	rows, err = Db.Queryx(qry2, LastIDEmpresa)

	if err != nil {
		log.Fatal(err)
		code = 500
		msgreturn = "Erro MySQL:" + err.Error()
	}
	for rows.Next() {
		err = rows.StructScan(&AccountView)
	}
	defer rows.Close()

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
		c.JSON(200, AccountView)
	}
}

func AccountUpdate(c *gin.Context) {
	token := services.TokenBearer(c.GetHeader("Authorization"))
	AccountReceiveUpdate := c.Param("account") // RECEBE ID PELA URL

	code := 200
	msgreturn := ""
	qryinser := ""

	var p models.AccountCreate
	err := c.ShouldBindJSON(&p)

	p.Password = services.SHA256Encoder(p.Password)

	Db := databases.GetDatabase()

	if AccountReceiveUpdate == "" {
		qryinser = "UPDATE accounts SET password = ? WHERE token = ?"
		_, err = Db.Queryx(qryinser, p.Password, token)
	} else {
		qryinser = "UPDATE accounts SET password = ? WHERE account_number = ?"
		_, err = Db.Queryx(qryinser, p.Password, AccountReceiveUpdate)
	}

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
		c.JSON(200, "Account Updated")
	}
}

func AccountDelete(c *gin.Context) {
	AccountReceive := c.Param("account") // RECEBE ID PELA URL
	token := services.TokenBearer(c.GetHeader("Authorization"))

	code := 200
	msgreturn := ""

	LiberaDeletar := false
	Db := databases.GetDatabase()

	/*valida se o token nao é a conta que esta sendo deletada*/

	AccountNumber := models.AccountNumber{}

	qry := "SELECT a.account_number FROM accounts a WHERE a.token = ? LIMIT 1"
	rows, err := Db.Queryx(qry, token)

	if err != nil {
		log.Fatal(err)
		code = 500
		msgreturn = "Erro MySQL:" + err.Error()
	}
	for rows.Next() {
		err = rows.StructScan(&AccountNumber)
	}

	if err != nil {
		log.Fatal(err)
		code = 500
		msgreturn = "Erro Leitura Linhas MySQL:" + err.Error()
	}
	defer rows.Close()

	if strconv.Itoa(AccountNumber.Account_number) != AccountReceive {
		LiberaDeletar = true
	}
	fmt.Println("DBT:", AccountNumber.Account_number)
	fmt.Println("token:", token)
	/*valida se o token nao é a conta que esta sendo deletada*/

	if LiberaDeletar == true {
		qry := "DELETE FROM accounts WHERE account_number = ?"
		_, err := Db.Queryx(qry, AccountReceive)

		if err != nil {
			log.Fatal(err)
			code = 500
			msgreturn = "Erro MySQL:" + err.Error()
		}
	} else {
		code = 401
		msgreturn = "you are not allowed to self-delete"
	}

	if code != 200 {
		c.JSON(code, gin.H{
			"error": msgreturn,
		})
		return
	} else {
		c.JSON(200, "Account Deleted")
	}
}
