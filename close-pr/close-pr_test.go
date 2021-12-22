package main

import (
	"os"
	"testing"
)

func TestPullRequestClose(t *testing.T) {
	type args struct {
		token      string
		githubrepo string
		githubref  string
	}
	tests := []struct {
		name        string
		args        args
		wantErr     bool
		wantSuccess bool
	}{
		{
			name: "Successful Close PR",
			args: args{
				token:      os.Getenv("GITHUB_OAUTH_TOKEN"),
				githubrepo: "jackstockley89/golangwebpage",
				githubref:  "refs/pull/120/merge",
			},
			wantErr:     false,
			wantSuccess: true,
		},
		{
			name: "Fail Close PR",
			args: args{
				token:      os.Getenv("GITHUB_OAUTH_TOKEN"),
				githubrepo: "%/%",
				githubref:  "refs/pull/000/merge",
			},
			wantErr:     true,
			wantSuccess: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PullRequestClose(tt.args.token, tt.args.githubrepo, tt.args.githubref)
		})
	}
}
