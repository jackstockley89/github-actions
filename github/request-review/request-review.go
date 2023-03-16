package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/google/go-github/github"
	"github.com/jackstockley89/github-actions/github/lib"
)

var (
	token      = flag.String("token", os.Getenv("GITHUB_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref  = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
	c          = lib.ClientConnect(*token)
	g          = lib.GetPullRequestData(*githubrepo, *githubref, *token)
	col        = lib.GetCollaborators(g.Owner, g.Repository, *token)
)

func assignReviewers() {
	// assign reviewers if they are collaborators and not the PR author
	var reviewers []string
	for _, collaborator := range col.Collaborators {
		if collaborator != g.User {
			reviewers = append(reviewers, collaborator)
		}
	}
	// create a review request
	reviewersRequest := &github.ReviewersRequest{
		Reviewers: reviewers,
	}
	_, _, err := c.PullRequests.RequestReviewers(context.Background(), g.Owner, g.Repository, g.Number, *reviewersRequest)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	assignReviewers()
}
