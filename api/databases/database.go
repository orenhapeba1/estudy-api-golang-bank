package databases

import (
	"fmt"
	"github.com/jmoiron/sqlx/reflectx"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

// OpenConn - Recebe dados da conexão em string e devolve a conexão postgres testada
func OpenConn() {
	fmt.Println("\nConectando ao MySQL...")
	//dbctx := context.Background()
	db, err := sqlx.Connect("mysql", "root:root@tcp(mysql:3306)/db?parseTime=true")
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conexão com MySQL efetuada e testada com sucesso!")
	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(10000)
	db.SetConnMaxLifetime(time.Hour)
	Db = db

	/*teste de models com DBS*/
	Db.Mapper = reflectx.NewMapperTagFunc("db",
		nil,
		func(s string) string {
			fmt.Println("39", s)
			return strings.ToLower(s)
		},
	)
	/*teste de models com DBS*/

}

func GetDatabase() *sqlx.DB {
	return Db
}
