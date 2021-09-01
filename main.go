package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

//get repo owner
func repoOwner() {

	token := os.Getenv("GITHUB_OAUTH_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list collaborators under "golangwebpage" repo
	repos, _, err := client.Repositories.ListCollaborators(context.Background(), "jackstockley89", "golangwebpage", nil)
	if err != nil {
		log.Fatalln(err)
	}
	for _, collab := range repos {
		fmt.Println("Collaborator:", *collab.Login)
	}
}

//get pr owner
func prOwner() {

	token := os.Getenv("GITHUB_OAUTH_TOKEN")

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list public pull request for "golangwebpage" repo
	prs, _, err := client.PullRequests.List(context.Background(), "jackstockley89", "golangwebpage", nil)
	if err != nil {
		log.Fatalln(err)
	}
	for _, pr := range prs {
		fmt.Println("Pull Request Title:", *pr.Title, "\nPull Request User:", *pr.User.Login)
	}
}

func main() {
	repoOwner()
	prOwner()
	//compare pr and repo owner

	//if pr and repo owner match check passes
	//if fails delete pr
}
