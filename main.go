package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

type Collaborators struct {
	User []string
}

type PullRequests struct {
	Title []string
	User  []string
}

// get repo owner
func getRepoOwner() {
	// get rnv token
	// Connect to giithub
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
		collabUserslice := Collaborators{
			User: []string{*collab.Login}}
		fmt.Println("Collaborator:", collabUserslice)
	}
}

// get pr owner
func getPrOwner() {
	// get rnv token
	// Connect to giithub
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
		prTitleslice := PullRequests{
			Title: []string{*pr.Title}}
		prUserslice := PullRequests{
			User: []string{*pr.User.Login}}
		fmt.Println("\nPull Request Title:", prTitleslice, "\nPull Request User:", prUserslice)
	}

}

// compare pr and repo owner
func compareCheck() {
	resp1 := "jackstockley89" // This should come from the getRepoOwner
	resp2 := "jackstockley89" // This should come from the getPrOwner
	// pull in data from struct for both collba and pr
	if resp1 == resp2 {
		fmt.Println("Match")
	} else {
		fmt.Println("No Match")
	}
}

func main() {
	getRepoOwner()
	getPrOwner()
	compareCheck()
	// if pr and repo owner match check passes
	// if fails delete pr
}
