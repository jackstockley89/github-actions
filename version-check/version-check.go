package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/v38/github"
	githubaction "github.com/sethvargo/go-githubactions"
	"golang.org/x/oauth2"
)

// PullRequestCheck will validate that the User of the pull request is a valid Collaborator
func versionCheck(token, githubrepo string) (bool, error) {
	// get env token
	// Connect to giithub
	var client *github.Client
	if token == "" {
		client = github.NewClient(nil)
	} else {
		ctx := context.Background()
		ts := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token},
		)
		tc := oauth2.NewClient(ctx, ts)

		client = github.NewClient(tc)
	}

	//repo user and repo name
	githubrepoS := strings.Split(githubrepo, "/")
	repoUser := githubrepoS[0]
	repoName := githubrepoS[1]

	t, _, err := client.Repositories.GetLatestRelease(context.Background(), repoUser, repoName)
	if err != nil {
		return false, err
	}

	ta := []string{*t.TagName}
	tj := strings.Join(ta, " ")
	output = fmt.Sprintln(tj)
	fmt.Println("Release Name: ", []string{*t.Name})
	fmt.Println("Release Tag: ", tj)

	return true, nil

}

var (
	token      = flag.String("token", os.Getenv("GITHUB_OAUTH_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	output     string
)

func main() {
	flag.Parse()
	t, err := versionCheck(*token, *githubrepo)
	if err != nil {
		log.Fatal(err)
	}
	if t {
		githubaction.SetOutput("Release Tag: ", output)
	}
}
