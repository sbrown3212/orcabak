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
	ErrInvalidGitStatusRenameLine      = errors.New("invalid git status rename line")
	ErrIncorrectAmountOfPartsToLine    = errors.New("incorrect amount of parts to line")
	ErrNoOrigPath                      = errors.New("renamed or copied FileStatus is missing OrigPath")
	ErrRenameSplitPaths                = errors.New("unable to split rename paths")
	ErrUnrecognizedGitStatusLinePrefix = errors.New("prefix not recognized for git status ouput")
)

func parseGitStatus(data []byte) (GitStatus, error) {
	result := GitStatus{}
	scanner := bufio.NewScanner(bytes.NewReader(data))
	for scanner.Scan() {
		line := scanner.Bytes()

		if len(line) == 0 {
			continue
		}

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
			err := parseRenamedCopiedLine(line, &result)
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
			return GitStatus{}, ErrUnrecognizedGitStatusLinePrefix
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
	parts := bytes.SplitN(line, []byte(" "), 9)
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

func parseRenamedCopiedLine(line []byte, g *GitStatus) error {
	// Rename and Copied output has a tab separating new path and original path
	// First, split by tab
	tabSplitParts := bytes.SplitN(line, []byte("\t"), 2)
	if len(tabSplitParts) != 2 {
		return ErrRenameSplitPaths
	}
	origPath := tabSplitParts[1]
	left := tabSplitParts[0]

	// Second, split by space
	parts := bytes.SplitN(left, []byte(" "), 10)
	if len(parts) < 10 {
		return ErrIncorrectAmountOfPartsToLine
	}

	newPath := parts[9]

	statusType := getStatusType(line[2])

	newFileStatus := FileStatus{
		Path:     string(newPath),
		OrigPath: string(origPath),
		Type:     statusType,
	}

	if newFileStatus.OrigPath == "" {
		return ErrNoOrigPath
	}

	g.Staged = append(g.Staged, newFileStatus)

	return nil
}

func parseConflicts(line []byte, g *GitStatus) error {
	parts := bytes.SplitN(line, []byte(" "), 11)
	if len(parts) != 11 {
		return ErrIncorrectAmountOfPartsToLine
	}

	path := parts[10]

	g.Conflicts = append(g.Conflicts, string(path))

	return nil
}

func parseUntracked(line []byte, g *GitStatus) error {
	parts := bytes.SplitN(line, []byte(" "), 2)
	if len(parts) != 2 {
		return ErrIncorrectAmountOfPartsToLine
	}

	path := parts[1]

	g.Untracked = append(g.Untracked, string(path))

	return nil
}
