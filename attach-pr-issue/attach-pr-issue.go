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
	githubaction "github.com/sethvargo/go-githubactions"
	"golang.org/x/oauth2"
)

func GetIssueNumber(token, githubrepo, githubref string) bool {
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
		log.Fatal()
	}
	issue := []string{*prs.Body}
	issueJ := strings.Join(issue, "\n")
	issueF := strings.Fields(issueJ)
	issueT := issueF[0]
	numberS := issueF[1]

	outcome := strings.Contains(issueT, "Issue-Number:")
	number, err := strconv.Atoi(numberS)
	if err != nil {
		return false
	}

	issuen = number
	fmt.Println(outcome, issuen)
	return outcome
}

func FindIssueNumber(token, githubrepo string) {
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

	issues, _, err := client.Issues.Get(context.Background(), repoUser, repoName, issuen)
	if err != nil {
		log.Fatal()
	}
	issuenumber := []int{*issues.Number}
	issueid := []int64{*issues.ID}
	issuet = []string{*issues.Title}
	issueurl = []string{*issues.URL}
	json, _ := json.MarshalIndent(issues, "", "  ")
	fmt.Println(string(json))
	fmt.Println("\nIssue ID:", issueid, "\nIssue Number:", issuenumber, "\nTitle:", issuet, "\nIssueURL:", issueurl)
	if err != nil {
		log.Fatal()
	}
}

func UpdatePRLink(token, githubrepo, githubref string) {
	// get env token
	//Connect to giithub
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

	// repo user and repo name
	githubrepoS := strings.Split(githubrepo, "/")
	repoUser := githubrepoS[0]
	repoName := githubrepoS[1]

	// get pr owner
	//	githubrefS := strings.Split(githubref, "/")
	//	branch := githubrefS[2]
	//	bid, _ := strconv.Atoi(branch)
	//
	//	issueurlJ := strings.Join(issueurl, "")
	//	issueURL := issueurlJ
	//	event := "closes"

	prlink, _, err := client.Issues.Get(context.Background(), repoUser, repoName, 127)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("\nChange:", prlink.Labels)
}

var (
	token      = flag.String("token", os.Getenv("GITHUB_OAUTH_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref  = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
	issuen     int
	issuet     []string
	issueurl   []string
)

func main() {
	flag.Parse()
	t := GetIssueNumber(*token, *githubrepo, *githubref)
	if t == true {
		githubaction.SetOutput("IssueFound", "true")
		FindIssueNumber(*token, *githubrepo)
		UpdatePRLink(*token, *githubrepo, *githubref)
	}
	if t == false {
		githubaction.SetOutput("IssueFound", "false")
	}
}
