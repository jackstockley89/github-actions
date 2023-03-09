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
func CreateComment(token, owner, repository, body string, bid int) {
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
func CreateReview(token, owner, repository, body string, bid int) (bool, error) {
	//create review comment on pull request
	review := &github.PullRequestReviewRequest{
		Body:  github.String(body),
		Event: github.String("COMMENT"),
	}
	c.PullRequests.CreateReview(context.Background(), owner, repository, bid, review)
	return true, nil
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
