package git

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/google/go-cmp/cmp"
)

const testdataDir = "testdata"

type testCase struct {
	input    string
	expected GitStatus
	wantErr  bool
}

// TODO: create cases for:
// - [ ]
// - [x] clean repository
// - [x] untracked file
// - [ ] modified_index file
// - [x] modified_worktree file
// - [ ] added file
// - [ ] deleted file
// - [ ] renamed file
// - [ ] conflict
// - [ ] staged and modified

// TODO: create cases for: (optionally)
// - [ ] branch ahead
// - [ ] branch behinc
// - [ ] branch ahead and behind

func TestParseGitStatus(t *testing.T) {
	cases := map[string]testCase{
		"untracked": {
			input: "untracked.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				// Untracked: ["hello.txt"],
				Untracked: []string{"hello.txt"},
			},
			wantErr: false,
		},
		"modified_worktree": {
			input: "modified.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Unstaged: []FileStatus{fs("hello.txt", "", StatusModified)},
			},
			wantErr: false,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			// read input file
			path := filepath.Join(testdataDir, tc.input)
			dat, err := os.ReadFile(path)
			if err != nil {
				t.Fatalf("test: %v - could not read input file: %v", name, err)
			}

			// parse input file
			output, err := parseGitStatus(dat)
			if err != nil {
				t.Fatalf("test: %v - git status parse failed: %v", name, err)
			}

			// compare parsed output with expected output
			if diff := cmp.Diff(tc.expected, output); diff != "" {
				t.Errorf("test: %v - mismatch (-expected, +actual):\n%v", name, diff)
			}
		})
	}
}

func fs(path, origPath string, status StatusType) FileStatus {
	result := FileStatus{
		Path: path,
		Type: status,
	}

	if origPath != "" {
		result.OrigPath = origPath
	}

	return result
}
