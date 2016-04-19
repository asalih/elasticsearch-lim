package main

import (
	"log"

	"github.com/asalih/elasticsearch-lim/app"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var timeh = app.TimeHandler{}
	go timeh.InitTime()

	app.InitServer()

}
