package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/google/go-github/github"
	"github.com/jackstockley89/github-actions/github-api/client"
	pullrequestinfo "github.com/jackstockley89/github-actions/github-api/pull-request-info"
)

var (
	token      = flag.String("token", os.Getenv("GITHUB_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref  = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
	c          = client.ClientConnect(*token)
	pri        = pullrequestinfo.PullRequestData(*githubrepo, *githubref)
	prauth     string
)

type Collaborators []struct {
	Login string `json:"login"`
}

type Reviewer []struct {
	Login string `json:"login"`
}

func PullRequestsAuth() {
	prs, _, err := c.PullRequests.Get(context.Background(), pri.Owner, pri.Repository, pri.Bid)
	if err != nil {
		log.Fatal(err)
	}
	prauth = prs.User.GetLogin()
}

func AssignUser() {
	// get Collaborators List for the given Respository
	// As this Respository has not teams it has collect the users name to pass into the review request
	var col Collaborators
	h, _, err := c.Repositories.ListCollaborators(context.Background(), pri.Owner, pri.Repository, &github.ListCollaboratorsOptions{})
	if err != nil {
		log.Fatal(err)
	}
	j, _ := json.Marshal(h)
	u := json.Unmarshal(j, &col)
	if u != nil {
		log.Panicln(u)
	}

	// Taken the List of Collaborators from CollectCollabList function and breaking it down into []string that is readable by the ReviewerRequest
	cprint := fmt.Sprintln(col)
	creplace := strings.ReplaceAll(cprint, "[", "")
	creplace2 := strings.ReplaceAll(creplace, "{", "")
	creplace3 := strings.ReplaceAll(creplace2, "}", "")
	creplace4 := strings.ReplaceAll(creplace3, "]", "")
	csplit := strings.Split(creplace4, " ")

	prAuth := prauth
	var r Reviewer
	if prAuth != csplit[0] {
		i := github.ReviewersRequest{Reviewers: []string{csplit[0]}}
		prs, _, err := c.PullRequests.RequestReviewers(context.Background(), pri.Owner, pri.Repository, pri.Bid, i)
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
		prs, _, err := c.PullRequests.RequestReviewers(context.Background(), pri.Owner, pri.Repository, pri.Bid, i)
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

func main() {
	flag.Parse()
	PullRequestsAuth()
	AssignUser()
}
