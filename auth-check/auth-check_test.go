package main

import (
	"testing"
)

func Test_pullRequestCheck(t *testing.T) {
	tests := []struct {
		name    string
		want    string
		wantErr bool
	}{
		{
			name:    "Test PullRequestCheck Function - Valid User",
			want:    "jackstockley89",
			wantErr: true,
		},
		{
			name:    "Test PullRequestCheck Function - Invalid User",
			want:    "testuser",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := pullRequestCheck()
			if (err != nil) != tt.wantErr {
				t.Errorf("PullRequestCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("PullRequestCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_collaboratorCheck(t *testing.T) {
	type args struct {
		pr string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Test CollaboratorCheck Function - Valid User",
			args: args{
				pr: "jackstockley89",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "Test CollaboratorCheck Function - Invalid User",
			args: args{
				pr: "testuser",
			},
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := collaboratorCheck(tt.args.pr)
			if (err != nil) != tt.wantErr {
				t.Errorf("CollaboratorCheck() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CollaboratorCheck() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pullRequestComment(t *testing.T) {
	type args struct {
		collab bool
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test pullRequestComment Function - Valid User",
			args: args{
				collab: true,
			},
		},
		{
			name: "Test pullRequestComment Function - Invalid User",
			args: args{
				collab: false,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pullRequestComment(tt.args.collab)
		})
	}
}
