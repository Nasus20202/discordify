package spotify

const (
	port      = ":8888"
	endpoint  = "/callback"
	state     = "discordify"
	cacheFile = ".refresh_token"

	spotifyIDEnv     = "SPOTIFY_ID"
	spotifySecretEnv = "SPOTIFY_SECRET"
)

var redirectURI = "http://localhost" + port + endpoint
