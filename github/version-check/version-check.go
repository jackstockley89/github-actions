package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	githubaction "github.com/sethvargo/go-githubactions"

	"github.com/jackstockley89/github-actions/github/lib"
)

var (
	token      = flag.String("token", os.Getenv("GITHUB_TOKEN"), "GihHub Personel token string")
	githubrepo = flag.String("githubrepo", os.Getenv("GITHUB_REPOSITORY"), "Github Repository string")
	githubref  = flag.String("githubref", os.Getenv("GITHUB_REF"), "Github Reference string")
	output     string
	c          = lib.ClientConnect(*token)
	pri        = lib.PullRequestData(*githubrepo, *githubref)
)

// versionCheck will check the latest release of a github repository
func versionCheck(token, githubrepo string) (bool, error) {
	t, _, err := c.Repositories.GetLatestRelease(context.Background(), pri.Owner, pri.Repository)
	if err != nil {
		return false, err
	}

	ta := []string{*t.TagName}
	tj := strings.Join(ta, " ")
	output = fmt.Sprintln(tj)
	fmt.Println("Release Name: ", []string{*t.Name})
	fmt.Println("Release Tag: ", tj)

	return true, nil

}

func main() {
	flag.Parse()
	t, err := versionCheck(*token, *githubrepo)
	if err != nil {
		log.Fatal(err)
	}
	if t {
		githubaction.SetOutput("release_tag", output)
	}
}
