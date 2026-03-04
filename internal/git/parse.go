package git

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	ErrUnexpectedFormatGitStatusOutput = errors.New("unexpected format for git status output")
	ErrInvalidGitStatusRenameLine      = errors.New("invalid git status rename line")
	ErrIncorrectAmountOfPartsToLine    = errors.New("incorrect amount of parts to line")
	ErrRenameSplitPaths                = errors.New("unable to split rename paths")
)

func parseGitStatus(data []byte) (GitStatus, error) {
	result := GitStatus{}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Bytes()

		switch line[0] {
		case '#': // Parse Branch lines
			err := parseBranch(line, &result)
			if err != nil {
				return GitStatus{}, fmt.Errorf("git status error: %w", err)
			}

		case '1': // Ordinary tracked line
			err := parseOrdinaryTracked(line, &result)
			if err != nil {
				return GitStatus{}, fmt.Errorf("error parsing ordinary tracked file line: %w", err)
			}

		case '2': // Renamed or copied lines
			err := parseRenameLine(line, &result)
			if err != nil {
				return GitStatus{}, fmt.Errorf("error parsing renamed file line: %w", err)
			}

		case 'u': // Files with conflicts
			err := parseConflicts(line, &result)
			if err != nil {
				return GitStatus{}, fmt.Errorf("error parsing conflict file line: %w", err)
			}

		case '?': // Untracked files
			err := parseUntracked(line, &result)
			if err != nil {
				return GitStatus{}, fmt.Errorf("error parsing untracked file line: %w", err)
			}

		default:
			return GitStatus{}, ErrUnexpectedFormatGitStatusOutput
		}
	}

	return result, nil
}

func parseBranch(line []byte, g *GitStatus) error {
	parts := bytes.Split(line, []byte(" "))

	switch {
	case bytes.HasPrefix(line, []byte("# branch.head")):
		g.Branch.Name = string(parts[2])

	case bytes.HasPrefix(line, []byte("# branch.upstream")):
		g.Branch.Upstream = string(parts[2])

	case bytes.HasPrefix(line, []byte("# branch.ab")):
		aheadStr := strings.TrimPrefix(string(parts[2]), "+")
		ahead, err := strconv.Atoi(aheadStr)
		if err != nil {
			return fmt.Errorf("error parsing git status branch line: %w", err)
		}
		g.Branch.Ahead = ahead

		behindStr := strings.TrimPrefix(string(parts[3]), "-")
		behind, err := strconv.Atoi(behindStr)
		if err != nil {
			return fmt.Errorf("error parsing git status branch line: %w", err)
		}
		g.Branch.Behind = behind
	}

	return nil
}

func getStatusType(char byte) StatusType {
	switch char {
	case 'M':
		return StatusModified
	case 'A':
		return StatusAdded
	case 'D':
		return StatusDeleted
	case 'R':
		return StatusRenamed
	case 'C':
		return StatusCopied
	case 'T':
		return StatusTypeChange
	}
	return StatusNone
}

func parseOrdinaryTracked(line []byte, g *GitStatus) error {
	parts := bytes.Split(line, []byte(" "))
	if len(parts) != 9 {
		return ErrIncorrectAmountOfPartsToLine
	}

	code := parts[1]
	path := string(parts[8])

	stagedStatus := getStatusType(code[0])
	unstagedStatus := getStatusType(code[1])

	if stagedStatus != StatusNone {
		g.Staged = append(g.Staged, FileStatus{
			Path: path,
			Type: stagedStatus,
		})
	}
	if unstagedStatus != StatusNone {
		g.Unstaged = append(g.Unstaged, FileStatus{
			Path: path,
			Type: unstagedStatus,
		})
	}

	return nil
}

func parseRenameLine(line []byte, g *GitStatus) error {
	// First, split by tab
	tabSplitParts := bytes.SplitN(line, []byte("\t"), 2)

	if len(tabSplitParts) != 2 {
		return ErrRenameSplitPaths
	}
	origPath := tabSplitParts[1]
	left := tabSplitParts[0]

	// Second, split by space
	parts := bytes.Split(left, []byte(" "))
	if len(parts) < 10 {
		return ErrIncorrectAmountOfPartsToLine
	}

	newPath := parts[9]

	statusType := getStatusType(line[0])

	g.Staged = append(g.Staged, FileStatus{
		Path:     string(newPath),
		OrigPath: string(origPath),
		Type:     statusType,
	})

	return nil
}

func parseConflicts(line []byte, g *GitStatus) error {
	parts := bytes.Split(line, []byte(" "))
	if len(parts) != 10 {
		return ErrIncorrectAmountOfPartsToLine
	}

	path := parts[9]

	g.Conflicts = append(g.Conflicts, string(path))

	return nil
}

func parseUntracked(line []byte, g *GitStatus) error {
	parts := bytes.Split(line, []byte(" "))
	if len(parts) != 2 {
		return ErrIncorrectAmountOfPartsToLine
	}

	path := parts[1]

	g.Untracked = append(g.Untracked, string(path))

	return nil
}
