package get

import (
	"context"
	"log"

	client "github.com/jackstockley89/github-actions/github-api/client"
)

func getCommitId(owner, repository, token string, bid int) string {
	c := client.ClientConnect(token)
	// get commit id for pull request
	prs, _, err := c.PullRequests.Get(context.Background(), owner, repository, bid)
	if err != nil {
		log.Fatal(err)
	}
	title := *prs.Title
	state := *prs.State
	url := *prs.URL
	commitID := *prs.Head.SHA

	log.Println("Pull Request Title:", title)
	log.Println("Pull Request State:", state)
	log.Println("Pull Request URL:", url)
	log.Println("Pull Request Commit ID:", commitID)

	return commitID
}
