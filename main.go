package main

import (
	"log"
	"netsepio-api/app"
)

func main() {
	app.Init()
	err := app.GinApp.Run(":8000")
	if err != nil {
		log.Fatal(err)
	}
}
