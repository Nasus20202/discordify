package spotify

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"golang.org/x/oauth2"
)

var (
	ch     = make(chan *spotify.Client)
	auth   *spotifyauth.Authenticator
	url    string
	client *spotify.Client
)

func GetClient(ctx context.Context) (*spotify.Client, error) {
	if client != nil {
		return client, nil
	}

	if os.Getenv(spotifyIDEnv) == "" || os.Getenv(spotifySecretEnv) == "" {
		return nil, fmt.Errorf("missing %s or %s environment variable", spotifyIDEnv, spotifySecretEnv)
	}

	auth = spotifyauth.New(spotifyauth.WithRedirectURL(redirectURI), spotifyauth.WithScopes(spotifyauth.ScopeUserReadCurrentlyPlaying, spotifyauth.ScopeUserReadPlaybackState))

	if _, err := os.Stat(cacheFile); err == nil {
		refreshToken, err := os.ReadFile(cacheFile)
		if err != nil {
			return nil, err
		}

		client := spotify.New(auth.Client(ctx, &oauth2.Token{RefreshToken: string(refreshToken)}))
		return client, nil
	}

	srv := &http.Server{Addr: port}
	http.HandleFunc(endpoint, completeAuth)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, url, http.StatusFound)
	})

	go func() {
		log.Printf("Server is running on localhost%s\n", port)
		if err := srv.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
		log.Print("Server is shutting down...")
	}()

	url = auth.AuthURL(state)
	log.Print("Please log in to Spotify by visiting the following page in your browser:", url)

	client := <-ch

	if err := srv.Shutdown(ctx); err != nil {
		return nil, err
	}

	token, err := client.Token()
	if err != nil {
		return nil, err
	}

	if err := os.WriteFile(cacheFile, []byte(token.RefreshToken), 0644); err != nil {
		return nil, err
	}

	return client, nil
}

func completeAuth(w http.ResponseWriter, r *http.Request) {
	tok, err := auth.Token(r.Context(), state, r)
	if err != nil {
		http.Error(w, "Couldn't get token", http.StatusForbidden)
		log.Println(err)
		return
	}
	if st := r.FormValue("state"); st != state {
		http.NotFound(w, r)
		log.Printf("State mismatch: %s != %s\n", st, state)
		return
	}

	client := spotify.New(auth.Client(r.Context(), tok))
	w.Write([]byte("Login Completed!"))
	ch <- client
}
