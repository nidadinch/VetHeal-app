package main

import (
	"log"
	"seniorproject-backend/server"
)

func main() {

	err := server.NewServer().StartServer()
	if err != nil {
		log.Fatalln(err)
	}
}
