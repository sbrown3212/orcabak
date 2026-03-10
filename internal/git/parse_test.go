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

func TestParseGitStatus(t *testing.T) {
	cases := map[string]testCase{
		"clean": {
			input: "clean.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
			},
			wantErr: false,
		},
		"untracked": {
			input: "untracked.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Untracked: []string{"hello.txt"},
			},
			wantErr: false,
		},
		"modified_index": {
			input: "modified_index.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Staged: []FileStatus{fs("second.txt", "", StatusModified)},
			},
			wantErr: false,
		},
		"modified_worktree": {
			input: "modified_worktree.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Unstaged: []FileStatus{fs("hello.txt", "", StatusModified)},
			},
			wantErr: false,
		},
		"modified_index_and_worktree": {
			input: "modified_index_and_worktree.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Staged:   []FileStatus{fs("second.txt", "", StatusModified)},
				Unstaged: []FileStatus{fs("second.txt", "", StatusModified)},
			},
			wantErr: false,
		},
		"added": {
			input: "added.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Staged: []FileStatus{fs("fourth.txt", "", StatusAdded)},
			},
			wantErr: false,
		},
		"deleted_worktree": {
			input: "deleted_worktree.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Unstaged: []FileStatus{fs("fourth.txt", "", StatusDeleted)},
			},
			wantErr: false,
		},
		"deleted_index": {
			input: "deleted_index.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Staged: []FileStatus{fs("fourth.txt", "", StatusDeleted)},
			},
			wantErr: false,
		},
		"renamed": {
			input: "renamed.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Staged: []FileStatus{fs("hello_world.txt", "hello.txt", StatusRenamed)},
			},
			wantErr: false,
		},
		"conflict": {
			input: "conflict.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Conflicts: []string{"hello.txt"},
			},
			wantErr: false,
		},
		"multiline": {
			input: "multiline.txt",
			expected: GitStatus{
				Branch: BranchInfo{
					Name: "main",
				},
				Staged: []FileStatus{
					fs("hello_world.txt", "hello.txt", StatusRenamed),
					fs("second.txt", "", StatusAdded),
				},
				Untracked: []string{"third.txt"},
			},
			wantErr: false,
		},
		"prefix_error": {
			input:    "prefix_error.txt",
			expected: GitStatus{},
			wantErr:  true,
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
