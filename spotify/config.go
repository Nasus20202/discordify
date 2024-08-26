package spotify

import (
	"flag"
)

const (
	protocolDefault  = "http://"
	hostDefault      = "localhost"
	portDefault      = ":8888"
	endpointDefault  = "/callback"
	stateDefault     = "discordify"
	cacheFileDefault = ".refresh_token"

	spotifyIDEnv     = "SPOTIFY_ID"
	spotifySecretEnv = "SPOTIFY_SECRET"
)

var (
	redirectURI string

	protocol  string
	host      string
	port      string
	endpoint  string
	state     string
	cacheFile string
)

func init() {
	flag.StringVar(&cacheFile, "cache", cacheFileDefault, "path to the file where the refresh token will be stored")

	flag.StringVar(&protocol, "protocol", protocolDefault, "protocol for the local server")

	flag.StringVar(&host, "host", hostDefault, "host for the local server")

	flag.StringVar(&port, "port", portDefault, "port for the local server")

	flag.StringVar(&endpoint, "endpoint", endpointDefault, "endpoint for the local server")

	flag.StringVar(&state, "state", stateDefault, "state for the OAuth2 flow")

	redirectURI = protocol + host + port + endpoint
}
