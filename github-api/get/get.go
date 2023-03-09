package get

import (
	"context"
	"log"

	client "github.com/jackstockley89/github-actions/github-api/client"
	pullrequestinfo "github.com/jackstockley89/github-actions/github-api/pull-request-info"
)

var (
	pr *PullRequestOutputData
)

type PullRequestOutputData struct {
	ID int64
	Title,
	Body,
	State,
	URL,
	CommitID,
	User,
	Owner,
	Repository string
	Number,
	Commits int
}

// GetCommitID will get the commit id for a pull request
func GetPullRequestData(githubrepo, githubref, token string) *PullRequestOutputData {
	c := client.ClientConnect(token)
	pri := pullrequestinfo.PullRequestData(githubrepo, githubref)
	// get commit id for pull request
	prs, _, err := c.PullRequests.Get(context.Background(), pri.Owner, pri.Repository, pri.Bid)
	if err != nil {
		log.Fatal(err)
	}

	// append to struct skip if nil
	pr = &PullRequestOutputData{
		ID:         prs.GetID(),
		Title:      prs.GetTitle(),
		Body:       prs.GetBody(),
		State:      prs.GetState(),
		URL:        prs.GetURL(),
		CommitID:   prs.GetHead().GetSHA(),
		User:       prs.GetUser().GetLogin(),
		Owner:      pri.Owner,
		Repository: pri.Repository,
		Number:     prs.GetNumber(),
		Commits:    prs.GetCommits(),
	}

	return pr
}
