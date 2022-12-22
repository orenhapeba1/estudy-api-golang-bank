package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/orenhapeba1/estudy-api-golang-bank/databases"
	"github.com/orenhapeba1/estudy-api-golang-bank/models"
	"github.com/orenhapeba1/estudy-api-golang-bank/services"
)

func Login(c *gin.Context) {
	Db := databases.GetDatabase()
	code := 200
	titlemsg := ""
	msgreturn := ""
	returnmsgcomplet := ""

	var p models.Login

	err := c.ShouldBindJSON(&p)

	if err != nil {
		code = 400
		titlemsg = "error"
		msgreturn = "cannot bind JSON: " + err.Error()

	}

	qry := "SELECT password FROM accounts WHERE account_number = ? LIMIT 1"
	rows, err := Db.Queryx(qry, p.Account)

	if err != nil {
		code = 500
		titlemsg = "error"
		msgreturn = "MySQL Error - : " + err.Error()
	}

	//user := models.Login{}
	Accounts_Password_Validation := models.Accounts_Password_Validation{}
	Accounts_Token := models.Accounts_Token{}

	for rows.Next() {
		err = rows.StructScan(&Accounts_Password_Validation)
	}

	if err != nil {
		code = 500
		titlemsg = "error"
		msgreturn = "MySQL Error - : " + err.Error()
	}

	defer rows.Close()

	if services.SHA256Encoder(p.Password) != Accounts_Password_Validation.Password {
		code = 401
		titlemsg = "error"
		msgreturn = "invalid credentials"
	} else {

		token, err := services.NewJWTService().GenerateToken((p.Account))

		if err != nil {
			code = 500
			titlemsg = "error"
			msgreturn = "Error generate token : " + err.Error()
			c.JSON(code, returnmsgcomplet)
			return
		}

		qryUpdateToken := "UPDATE accounts SET token = ? WHERE account_number = ?"
		_, err = Db.Queryx(qryUpdateToken, token, p.Account)

		Accounts_Token.Token = token
	}

	if code != 200 {
		returnmsgcomplet = titlemsg + ":" + msgreturn
		c.JSON(code, returnmsgcomplet)

	} else {
		c.JSON(200, Accounts_Token)

	}

}
