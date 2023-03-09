package create

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	client "github.com/jackstockley89/github-actions/github-api/client"
)

var (
	token string
	c     = client.ClientConnect(token)
)

// CreateComment will create a comment on the pull request
func CreateComment(owner, repository, token, body string, bid int) {
	// create comment on pull request
	comment := &github.PullRequestComment{
		Body: github.String(body),
	}
	_, _, err := c.PullRequests.CreateComment(context.Background(), owner, repository, bid, comment)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateReview will create a review on the pull request
func CreateReview(owner, repository, token, body string, bid int) {
	// get commit id
	commitID := getcommitid.getCommitId(owner, repository, token, bid)
	// create review comment on pull request
	review := &github.PullRequestReviewRequest{
		CommitID: commitID,
		Body:     github.String(body),
		Event:    github.String("COMMENT"),
	}
	_, _, err := c.PullRequests.CreateReview(context.Background(), owner, repository, bid, review)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateIssue will create an issue in a repository
func CreateIssue(owner, repository, title, body string) {
	// create issue
	issue := &github.IssueRequest{
		Title: github.String(title),
		Body:  github.String(body),
	}
	_, _, err := c.Issues.Create(context.Background(), owner, repository, issue)
	if err != nil {
		log.Fatal(err)
	}
}
