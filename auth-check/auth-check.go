package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/jackstockley89/github-actions/github-api/client"
	pullrequestinfo "github.com/jackstockley89/github-actions/github-api/pull-request-info"
	githubaction "github.com/sethvargo/go-githubactions"
)

// PullRequestCheck will validate that the User of the pull request is a valid Collaborator
func PullRequestCheck() (bool, error) {
	flag.Parse()

	token := flag.String("token", os.Getenv("GITHUB_TOKEN"), "GihHub Personel token string")
	githubrepo := flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref := flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
	c := client.ClientConnect(*token)
	pri := pullrequestinfo.PullRequestData(*githubrepo, *githubref)

	prs, _, err := c.PullRequests.Get(context.Background(), pri.Owner, pri.Repository, pri.Bid)
	if err != nil {
		return false, err
	}

	prarray := []string{*prs.User.Login}
	pr := strings.Join(prarray, " ")
	fmt.Println("Pull Request User:", pr)

	// compare pr user with the repo collaborators
	repos, _, err := c.Repositories.IsCollaborator(context.Background(), pri.Owner, pri.Repository, pr)
	if err != nil {
		return false, err
	}
	fmt.Println("Collaborator Status:", repos)
	return repos, nil
}

func main() {
	t, err := PullRequestCheck()
	if err != nil {
		log.Fatal(err)
	}
	if t {
		githubaction.SetOutput("approve_colab_user", "true")
	}
}
