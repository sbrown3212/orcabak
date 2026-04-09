package git

func (g *GitCLIClient) Commit(orcaProfileDir string, args ...string) (string, error) {
	commitArgs := []string{"commit"}
	for _, msg := range args {
		commitArgs = append(commitArgs, "--message")
		commitArgs = append(commitArgs, msg)
	}

	commitOutput, err := g.Runner.Run(orcaProfileDir, "git", commitArgs...)
	if err != nil {
		return string(commitOutput), err
	}

	return string(commitOutput), nil
}
