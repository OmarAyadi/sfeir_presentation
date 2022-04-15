package main

import (
	"github.com/joho/godotenv"
	"log"
	"sfeir/global"
	"sfeir/server"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	global.LazyEnvVariableInit()

	s, err := server.NewServer()
	if err != nil {
		log.Fatal(err)
	}

	err = s.Run()
	if err != nil {
		log.Fatal(err)
	}
}
