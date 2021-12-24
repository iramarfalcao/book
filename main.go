package main

import "github.com/iramarfalcao/book/server"

func main() {
	server := server.NewServer()
	server.Run()
}
