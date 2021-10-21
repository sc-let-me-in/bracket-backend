package service

import (
	"bracket-backend/models"
	"context"
	"github.com/zmb3/spotify/v2"
	spotifyauth "github.com/zmb3/spotify/v2/auth"
	"os"
	"golang.org/x/oauth2/clientcredentials"
)

// GetSong gets song information from spotify
func GetSong(id string) models.Song {

	// adapted from https://github.com/zmb3/spotify/blob/master/examples/search/search.go
	// TODO: move this somewhere else where it can be reusable/done once at startup
	ctx := context.Background()
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotifyauth.TokenURL,
	}
	spToken, _ := config.Token(ctx)
	httpClient := spotifyauth.New().Client(ctx, spToken)
	spotifyClient := spotify.New(httpClient)

	fullTrack, _ := spotifyClient.GetTrack(ctx, spotify.ID(id))

	return models.Song{
		SongID: fullTrack.ID.String(),
		Title: fullTrack.Name,
		PreviewClip: fullTrack.PreviewURL,
	}
}