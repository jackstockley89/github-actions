package client

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// clientConnect Takens a token and uses this to establish a connection to github
func ClientConnect(token string) *github.Client {

	// get env token
	// Connect to giithub
	if token == "" {
		log.Fatal("token missing")
	}
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client := github.NewClient(tc)
	return client
}
