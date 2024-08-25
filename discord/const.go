package discord

import "fmt"

const (
	discordTokenEnv = "DISCORD_TOKEN"
	discordAPI      = "https://discord.com/api/v9"
	userSettings    = "/users/@me/settings"
)

var (
	ErrTokenNotFound = fmt.Errorf("discord token not found in %s environment variable", discordTokenEnv)
	ErrorStatus      = fmt.Errorf("error setting status")
)
