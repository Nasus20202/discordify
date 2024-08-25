package main

import (
	"context"
	"log"
	"strings"
	"time"

	"github.com/nasus20202/discordify/discord"
	client "github.com/nasus20202/discordify/spotify"
	"github.com/zmb3/spotify/v2"
)

const (
	interval                = 500 * time.Millisecond
	statusTypeDurationTicks = 4
	emoji                   = "🎶"
)

func RunLoop(ctx context.Context) error {
	log.Println("Starting Discordify...")

	client, err := client.GetClient(ctx)
	if err != nil {
		return err
	}

	user, err := client.CurrentUser(ctx)
	if err != nil {
		return err
	}

	log.Printf("You are logged in as: %s", user.ID)
	discord.ClearStatus(ctx)

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	counter := 0
	for {
		<-ticker.C

		if err := tick(ctx, counter, client); err != nil {
			log.Println("Error:", err)
		}
		counter++
	}
}

func tick(ctx context.Context, counter int, client *spotify.Client) error {
	playing, err := client.PlayerCurrentlyPlaying(ctx)
	if err != nil {
		return err
	}

	if !playing.Playing {
		return discord.ClearStatus(ctx)
	}

	statusFuncs := []func(*spotify.CurrentlyPlaying) string{
		createSongString,
		createArtistString,
	}
	currentStatusFunc := statusFuncs[(counter/statusTypeDurationTicks)%len(statusFuncs)]

	return discord.SetStatus(ctx, currentStatusFunc(playing), emoji)
}

func createSongString(playing *spotify.CurrentlyPlaying) string {
	return playing.Item.Name
}

func createArtistString(playing *spotify.CurrentlyPlaying) string {
	var sb strings.Builder

	for i, artist := range playing.Item.Artists {
		if i > 0 {
			sb.WriteString(", ")
		}
		sb.WriteString(artist.Name)
	}

	return sb.String()
}
