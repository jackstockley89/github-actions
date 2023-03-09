package get

import (
	"context"
	"log"

	"github.com/google/go-github/github"
	client "github.com/jackstockley89/github-actions/github-api/client"
	pullrequestinfo "github.com/jackstockley89/github-actions/github-api/pull-request-info"
)

var (
	pr *PullRequestOutputData
)

// PullRequestOutputData is a struct to hold the data for a pull request to be used in other functions and packages
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

// Collaborators is a struct to hold the data for a repository collaborators to be used in other functions and packages
type Collaborators struct {
	Collaborators []string
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

// GetCollaborators will get a list of collaborators for a repository
func GetCollaborators(owner, repository, token string) *Collaborators {
	c := client.ClientConnect(token)
	col := &Collaborators{}
	// get repository collaborators
	options := &github.ListCollaboratorsOptions{
		ListOptions: github.ListOptions{PerPage: 10},
	}
	collaborators, _, err := c.Repositories.ListCollaborators(context.Background(), owner, repository, options)
	if err != nil {
		log.Fatal(err)
	}

	// assign collaborators to a struct
	for _, collaborator := range collaborators {
		col.Collaborators = append(col.Collaborators, *collaborator.Login)
	}
	return col
}
