package main

import (
	"testing"
)

func Test_collaboratorCheck(t *testing.T) {
	tests := []struct {
		name    string
		want    bool
		wantErr bool
	}{
		{
			name:    "Test CollaboratorCheck Function - Valid User",
			want:    true,
			wantErr: false,
		},
		{
			name:    "Test CollaboratorCheck Function - Invalid User",
			want:    false,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := collaboratorCheck()
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
