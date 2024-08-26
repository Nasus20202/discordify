package main

import (
	"context"
	"flag"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	flag.Parse()

	err := RunLoop(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}
