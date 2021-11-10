package main

import (
	"os"
	"testing"
)

func TestPullRequestCheck(t *testing.T) {
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
			name: "Successful Collaborator Check",
			args: args{
				token:      os.Getenv("GITHUB_OAUTH_TOKEN"),
				githubrepo: "jackstockley89/golangwebpage",
				githubref:  "refs/pull/101/merge",
			},
			wantErr:     false,
			wantSuccess: true,
		},
		{
			name: "Fail Collaborator Check",
			args: args{
				token:      os.Getenv("GITHUB_OAUTH_TOKEN"),
				githubrepo: "%/%",
				githubref:  "refs/pull/001/merge",
			},
			wantErr:     true,
			wantSuccess: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := PullRequestCheck(tt.args.token, tt.args.githubrepo, tt.args.githubref)
			if (err != nil) != tt.wantErr {
				t.Errorf("PullRequestCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
			if b != tt.wantSuccess {
				t.Errorf("PullRequestCheck() error = %v, wantSuccess %v", err, tt.wantSuccess)
			}
		})
	}
}
