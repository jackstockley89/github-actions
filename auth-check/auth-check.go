package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/v38/github"
	githubaction "github.com/sethvargo/go-githubactions"
	"golang.org/x/oauth2"
)

// PullRequestCheck will validate that the User of the pull request is a valid Collaborator
func PullRequestCheck(token, githubrepo, githubref string) (bool, error) {
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
	prs, _, err := client.PullRequests.Get(context.Background(), repoUser, repoName, bid)
	if err != nil {
		return false, err
	}
	prarray := []string{*prs.User.Login}
	pr := strings.Join(prarray, " ")
	fmt.Println("Pull Request User:", pr)

	// compare pr user with the repo collaborators
	repos, _, err := client.Repositories.IsCollaborator(context.Background(), repoUser, repoName, pr)
	if err != nil {
		return false, err
	}
	fmt.Println("Collaborator Status:", repos)
	return repos, nil
}

var (
	token      = flag.String("token", os.Getenv("GITHUB_OAUTH_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref  = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
)

func main() {
	flag.Parse()
	t, err := PullRequestCheck(*token, *githubrepo, *githubref)
	if err != nil {
		log.Fatal(err)
	}
	if t {
		githubaction.SetOutput("approve_colab_user", "true")
	}
}
