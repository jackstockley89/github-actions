package lib

import (
	"strconv"
	"strings"
)

// Return values from pullRequestInfo
type PullRequestInputData struct {
	Owner,
	Repository string
	Bid int
}

// TODO: gather pull request data
func PullRequestData(githubrepo, githubref string) PullRequestInputData {
	var r PullRequestInputData
	//repo user and repo name
	githubrepoS := strings.Split(githubrepo, "/")
	r.Owner = githubrepoS[0]
	r.Repository = githubrepoS[1]

	// get pr owner
	githubrefS := strings.Split(githubref, "/")
	branch := githubrefS[2]
	r.Bid, _ = strconv.Atoi(branch)

	return r
}
