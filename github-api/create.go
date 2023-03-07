package create

import (
	"context"
	"flag"
	"log"
	"os"

	c "github.com/jackstockley89/github-actions/github-api/client"
	pr "github.com/jackstockley89/github-actions/github-api/pullrequestinfo"

	"github.com/google/go-github/github"
)

var (
	token           = flag.String("token", os.Getenv("GITHUB_OAUTH_TOKEN"), "GihHub Personel token string")
	githubrepo      = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref       = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Respository PR ref string")
	client          = c.ClientConnect(*token)
	pullRequestInfo = pr.PullRequest(*githubrepo, *githubref)
)

// CreateComment will create a comment on the pull request
func CreateComment() {
	// connect to github
	// create comment
	comment := &github.PullRequestComment{
		Body: github.String("This is a test comment"),
	}
	_, _, err := client.PullRequests.CreateComment(context.Background(), pullRequestInfo.Owner, pullRequestInfo.Repository, pullRequestInfo.Bid, comment)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateReview will create a review on the pull request
func CreateReview() {
	// connect to github
	// create review
	review := &github.PullRequestReviewRequest{
		Body: github.String("This is a test review"),
	}
	_, _, err := client.PullRequests.CreateReview(context.Background(), pullRequestInfo.Owner, pullRequestInfo.Repository, pullRequestInfo.Bid, review)
	if err != nil {
		log.Fatal(err)
	}
}

// CreateIssue will create an issue in a repository
func CreateIssue() {
	// connect to github
	// create issue
	issue := &github.IssueRequest{
		Title: github.String("This is a test issue"),
	}
	_, _, err := client.Issues.Create(context.Background(), pullRequestInfo.Owner, pullRequestInfo.Repository, issue)
	if err != nil {
		log.Fatal(err)
	}
}
