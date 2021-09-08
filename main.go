package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/go-github/v38/github"
	"golang.org/x/oauth2"
)

// get repo owner
func gatherInput() {
	var c []string
	// get rnv token
	// Connect to giithub
	token := os.Getenv("GITHUB_OAUTH_TOKEN")
	githubref := os.Getenv("GITHUB_REF")
	sgr := strings.Split(githubref, "/")
	branch := sgr[2]
	bid, _ := strconv.Atoi(branch)

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	// list collaborators under "golangwebpage" repo
	repos, _, err := client.Repositories.ListCollaborators(context.Background(), "jackstockley89", "golangwebpage", nil)
	if err != nil {
		log.Fatalln(err)
	}
	for _, collab := range repos {
		fmt.Println(*collab.Login)
		c = []string{*collab.Login}
	}

	// get pr owner
	// list public pull request for "golangwebpage" repo
	prs, _, err := client.PullRequests.Get(context.Background(), "jackstockley89", "golangwebpage", bid)
	if err != nil {
		log.Fatalln(err)
	}
	p := []string{*prs.User.Login}
	fmt.Println("\n", compareCheck([]string{fmt.Sprintln(c)}, []string{fmt.Sprintln(p)}))

}

// compare pr and repo owner
func compareCheck(a1 []string, a2 []string) bool {
	// a1 // This should come from the client.Repositories.ListCollaborators //"jackstockley89, jasonBirchall"
	// a2 // This should come from the client.PullRequests.Get //"jackstockley89" or "jasonBirchall"
	fmt.Println("variable: a1")
	fmt.Println(a1)
	fmt.Println("variable: a2")
	fmt.Println(a2)
	return reflect.DeepEqual(a1, a2)
}

func main() {
	gatherInput()
	// if pr and repo owner match check passes
	// if fails delete pr
}
