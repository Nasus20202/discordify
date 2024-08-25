package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	err := RunLoop(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
