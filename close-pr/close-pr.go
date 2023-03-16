package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/github"
	client "github.com/jackstockley89/github-actions/github-api/client"
	get "github.com/jackstockley89/github-actions/github-api/get"
)

func PullRequestClose() {
	flag.Parse()

	token := flag.String("token", os.Getenv("GITHUB_TOKEN"), "GihHub Personel token string")
	githubrepo := flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref := flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
	c := client.ClientConnect(*token)
	g := get.GetPullRequestData(*githubrepo, *githubref, *token)
	state := &github.PullRequest{State: github.String("closed")}

	prs, _, err := c.PullRequests.Edit(context.Background(), g.Owner, g.Repository, g.Number, state)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(prs)
}

func main() {
	PullRequestClose()
}
