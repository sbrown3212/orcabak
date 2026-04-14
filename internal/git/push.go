package git

func (g *GitCLIClient) Push(orcaConfigDir string) (string, error) {
	output, err := g.Runner.Run(orcaConfigDir, "git", "push")
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}
