package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v40/github"
	"golang.org/x/oauth2"
)

func PullRequestClose(token, githubrepo, githubref string) {
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

	// get pr owner
	githubrefS := strings.Split(githubref, "/")
	branch := githubrefS[2]
	bid, _ := strconv.Atoi(branch)
	state := "closed"

	prs, _, err := client.PullRequests.Edit(context.Background(), repoUser, repoName, bid, &github.PullRequest{State: &state})
	if err != nil {
		log.Fatal()
	}
	fmt.Println(prs)
}

var (
	token      = flag.String("token", os.Getenv("GITHUB_OAUTH_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref  = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
)

func main() {
	flag.Parse()
	PullRequestClose(*token, *githubrepo, *githubref)
}
