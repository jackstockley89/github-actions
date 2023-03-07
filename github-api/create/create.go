package create

import (
	"context"
	"log"

	client "github.com/jackstockley89/github-actions/github-api/client"
	pullrequestinfo "github.com/jackstockley89/github-actions/github-api/pull-request-info"

	"github.com/google/go-github/github"
)

var (
	token      string
	githubrepo string
	githubref  string
	c          = client.ClientConnect(token)
	pri        = pullrequestinfo.PullRequestData(githubrepo, githubref)
)

// CreateComment will create a comment on the pull request
func CreateComment(token, githubrepo, githubref string) {
	// connect to github
	// create comment
	comment := &github.PullRequestComment{
		Body: github.String("This is a test comment"),
	}
	_, _, err := c.PullRequests.CreateComment(context.Background(), pri.Owner, pri.Repository, pri.Bid, comment)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateReview will create a review on the pull request
func CreateReview(token, githubrepo, githubref string) {
	// connect to github
	// create review
	review := &github.PullRequestReviewRequest{
		Body: github.String("This is a test review"),
	}
	_, _, err := c.PullRequests.CreateReview(context.Background(), pri.Owner, pri.Repository, pri.Bid, review)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateIssue will create an issue in a repository
func CreateIssue(token, githubrepo, githubref string) {
	// connect to github
	// create issue
	issue := &github.IssueRequest{
		Title: github.String("This is a test issue"),
	}
	_, _, err := c.Issues.Create(context.Background(), pri.Owner, pri.Repository, issue)
	if err != nil {
		log.Fatal(err)
	}
}
