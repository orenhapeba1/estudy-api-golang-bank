package main

import (
	"github.com/orenhapeba1/estudy-api-golang-bank/databases"
	"github.com/orenhapeba1/estudy-api-golang-bank/server"
)

func main() {
	databases.OpenConn()
	s := server.NewServer()
	s.Run()
}
