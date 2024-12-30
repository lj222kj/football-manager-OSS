package main

import (
	"log"

	"github.com/lj222kj/app"
)

func main() {
	if err := app.New(); err != nil {
		log.Fatalf("error %v", err)
	}
}
