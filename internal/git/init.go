package git

func (g *GitCLIClient) Init(orcaProfileDir string) (string, error) {
	output, err := g.Runner.Run(orcaProfileDir, "git", "init")
	if err != nil {
		return "", err
	}

	return string(output), nil
}
