package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/orenhapeba1/estudy-api-golang-bank/databases"
	"github.com/orenhapeba1/estudy-api-golang-bank/models"
	"github.com/orenhapeba1/estudy-api-golang-bank/services"
	"log"
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
