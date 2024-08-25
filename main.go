package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/nasus20202/discordify/discordify"
)

func main() {
	godotenv.Load()
	ctx := context.Background()

	client, err := discordify.GetClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	user, err := client.CurrentUser(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("You are logged in as: %s", user.ID)
}
