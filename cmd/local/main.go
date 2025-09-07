package main

import (
	"log"
	"os"

	"heloo-go/internal/app"
)

func main() {
	server, _, err := app.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	addr := os.Getenv("PORT")
	if addr == "" {
		addr = "8080"
	}
	server.Logger.Fatal(server.Start(":" + addr))
}
