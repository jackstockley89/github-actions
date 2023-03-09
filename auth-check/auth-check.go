package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	client "github.com/jackstockley89/github-actions/github-api/client"
	createcomment "github.com/jackstockley89/github-actions/github-api/create"
	pullrequestinfo "github.com/jackstockley89/github-actions/github-api/pull-request-info"
)

var (
	token      = flag.String("token", os.Getenv("GITHUB_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref  = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
	c          = client.ClientConnect(*token)
	pri        = pullrequestinfo.PullRequestData(*githubrepo, *githubref)
	body       string
)

// PullRequestCheck will validate that the User of the pull request is a valid Collaborator
func pullRequestCheck() (string, error) {
	prs, _, err := c.PullRequests.Get(context.Background(), pri.Owner, pri.Repository, pri.Bid)
	if err != nil {
		return "", err
	}

	prarray := []string{*prs.User.Login}
	pr := strings.Join(prarray, " ")
	fmt.Println("Pull Request User:", pr)

	return pr, nil
}

func collaboratorCheck(pr string) (bool, error) {
	// compare pr user with the repo collaborators
	collab, _, err := c.Repositories.IsCollaborator(context.Background(), pri.Owner, pri.Repository, pr)
	if err != nil {
		return false, err
	}
	body = fmt.Sprintln("Collaborator Status:", collab)
	return collab, nil
}

func pullRequestComment(collab bool) {
	if !collab {
		fmt.Println("User is not a collaborator")
		createcomment.CreateReview(pri.Owner, pri.Repository, *token, body, pri.Bid)
	} else {
		fmt.Println("User is a collaborator")
		createcomment.CreateReview(pri.Owner, pri.Repository, *token, body, pri.Bid)
	}
}

func main() {
	flag.Parse()
	pr, err := pullRequestCheck()
	if err != nil {
		log.Fatal(err)
	}
	collab, err := collaboratorCheck(pr)
	if err != nil {
		log.Fatal(err)
	}
	pullRequestComment(collab)
}
