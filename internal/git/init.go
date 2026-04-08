package git

func (g *GitCLIClient) Init(repoDir string) (string, error) {
	output, err := g.Runner.Run(repoDir, "git", "init")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
