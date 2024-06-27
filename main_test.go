package main

import (
	"errors"
	"testing"
)

func TestFindOddOccurs(t *testing.T) {
	var tests = []struct {
		name      string
		args      []int
		want      int
		wantError error
	}{
		{
			name:      "testCase1: [7]",
			args:      []int{7},
			want:      7,
			wantError: nil,
		},
		{
			name:      "testCase2: [0]",
			args:      []int{0},
			want:      0,
			wantError: nil,
		},
		{
			name:      "testCase3: [1,1,2] ",
			args:      []int{1, 1, 2},
			want:      2,
			wantError: nil,
		},
		{
			name:      "testCase4: [0,1,0,1,0] ",
			args:      []int{0, 1, 0, 1, 0},
			want:      0,
			wantError: nil,
		},
		{
			name:      "testCase5: [1,2,2,3,3,3,4,3,3,3,2,2,1] ",
			args:      []int{1, 2, 2, 3, 3, 3, 4, 3, 3, 3, 2, 2, 1},
			want:      4,
			wantError: nil,
		},
		{
			name:      "testCase6: no odd number found",
			args:      []int{2, 3, 4, 2, 3, 4},
			want:      -1,
			wantError: errors.New("no odd number found"),
		},
		{
			name:      "testCase3: empty array",
			args:      []int{},
			want:      -1,
			wantError: errors.New("no odd number found"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := findOddOccurs(tt.args)
			if got != tt.want || (err != nil && err.Error() != tt.wantError.Error()) {
				if got != tt.want {
					t.Errorf("findOddOccurs(%v) = %v, want: %v", tt.args, got, tt.want)
				}
				if err != nil && err.Error() != tt.wantError.Error() {
					t.Errorf("Unexpected error: %v, want: %v", err, tt.wantError)
				}
			}
		})
	}
}
