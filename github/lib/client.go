package lib

import (
	"context"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

var (
	ctx    = context.Background()
	client *github.Client
)

// clientConnect Takens a token and uses this to establish a connection to github
func ClientConnect(token string) *github.Client {
	var client *github.Client
	// get env token
	// Connect to giithub
	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)
	client = github.NewClient(tc)
	if token == "" {
		client = github.NewClient(nil)
	}
	return client
}
