package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

type Collaborators []struct {
	Login string `json:"login"`
}

type Reviewer []struct {
	Login string `json:"login"`
}

func CallClient(token string) *github.Client {
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
	return client
}

func PullRequestsAuth(githubrepo, githubref string) {
	client := CallClient(*token)
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
		log.Fatal(err)
	}
	prauth = prs.User.GetLogin()
}

func AssignUser(githubrepo, githubref string) {
	client := CallClient(*token)
	//repo user and repo name
	githubrepoS := strings.Split(githubrepo, "/")
	repoUser := githubrepoS[0]
	repoName := githubrepoS[1]

	// get pr owner
	githubrefS := strings.Split(githubref, "/")
	branch := githubrefS[2]
	bid, _ := strconv.Atoi(branch)

	// get Collaborators List for the given Respository
	// As this Respository has not teams it has collect the users name to pass into the review request
	var c Collaborators
	h, _, err := client.Repositories.ListCollaborators(context.Background(), repoUser, repoName, &github.ListCollaboratorsOptions{})
	if err != nil {
		log.Fatal(err)
	}
	j, _ := json.Marshal(h)
	u := json.Unmarshal(j, &c)
	if u != nil {
		log.Panicln(u)
	}

	// Taken the List of Collaborators from CollectCollabList function and breaking it down into []string that is readable by the ReviewerRequest
	cprint := fmt.Sprintln(c)
	creplace := strings.ReplaceAll(cprint, "[", "")
	creplace2 := strings.ReplaceAll(creplace, "{", "")
	creplace3 := strings.ReplaceAll(creplace2, "}", "")
	creplace4 := strings.ReplaceAll(creplace3, "]", "")
	csplit := strings.Split(creplace4, " ")

	prAuth := prauth
	var r Reviewer
	if prAuth != csplit[0] {
		i := github.ReviewersRequest{Reviewers: []string{csplit[0]}}
		prs, _, err := client.PullRequests.RequestReviewers(context.Background(), repoUser, repoName, bid, i)
		if err != nil {
			log.Fatal(err)
		}
		formatjson, _ := json.Marshal(prs.RequestedReviewers)
		formatstr := json.Unmarshal(formatjson, &r)
		if formatstr != nil {
			log.Panicln(formatstr)
		}
	}
	if prAuth != csplit[1] {
		i := github.ReviewersRequest{Reviewers: []string{csplit[1]}}
		prs, _, err := client.PullRequests.RequestReviewers(context.Background(), repoUser, repoName, bid, i)
		if err != nil {
			log.Fatal(err)
		}
		formatjson, _ := json.Marshal(prs.RequestedReviewers)
		formatstr := json.Unmarshal(formatjson, &r)
		if formatstr != nil {
			log.Panicln(formatstr)
		}
	}
	result := &r
	fmt.Println("Requested Reviewer:", result)
}

var (
	prauth     string
	token      = flag.String("token", os.Getenv("GITHUB_OAUTH_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref  = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
)

func main() {
	flag.Parse()
	PullRequestsAuth(*githubrepo, *githubref)
	AssignUser(*githubrepo, *githubref)
}
