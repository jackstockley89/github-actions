package main

import "testing"

func TestPullRequestCheck(t *testing.T) {
	type args struct {
		token      string
		githubrepo string
		githubref  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful Collaborator Check",
			args: args{
				token:      "",
				githubrepo: "jackstockley89/golangwebpage",
				githubref:  "refs/pull/101/merge",
			},
			wantErr: false,
		},
		{
			name: "Fail Collaborator Check",
			args: args{
				token:      "",
				githubrepo: "%/%",
				githubref:  "refs/pull/001/merge",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := PullRequestCheck(tt.args.token, tt.args.githubrepo, tt.args.githubref); (err != nil) != tt.wantErr {
				t.Errorf("PullRequestCheck() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
