package main

import (
	"context"
	"flag"
	"log"
	"os"

	"github.com/google/go-github/github"
	client "github.com/jackstockley89/github-actions/github-api/client"
	get "github.com/jackstockley89/github-actions/github-api/get"
)

var (
	token      = flag.String("token", os.Getenv("GITHUB_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref  = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
	c          = client.ClientConnect(*token)
	g          = get.GetPullRequestData(*githubrepo, *githubref, *token)
	col        Collaborators
)

type Collaborators struct {
	Collaborators []string
}

func getCollaborators() {
	// get repository collaborators
	options := &github.ListCollaboratorsOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	collaborators, _, err := c.Repositories.ListCollaborators(context.Background(), g.Owner, g.Repository, options)
	if err != nil {
		log.Fatal(err)
	}

	// assign collaborators to a struct
	for _, collaborator := range collaborators {
		col.Collaborators = append(col.Collaborators, *collaborator.Login)
	}
}

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
	getCollaborators()
	assignReviewers()
}
