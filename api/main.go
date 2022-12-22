package main

import (
	"github.com/orenhapeba1/estudy-api-golang-bank/databases"
	"github.com/orenhapeba1/estudy-api-golang-bank/server"
)

func main() {

	//fmt.Printf(services.SHA256Encoder("4501bgui"))

	databases.OpenConn()
	s := server.NewServer()
	s.Run()
}
