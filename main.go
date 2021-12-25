package main

import (
	"github.com/iramarfalcao/book/database"
	"github.com/iramarfalcao/book/server"
)

func main() {
	database.StartDB()

	server := server.NewServer()
	server.Run()
}
