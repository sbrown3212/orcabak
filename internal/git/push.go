package git

func (g *GitCLIClient) Push(orcaConfigDir string) (string, error) {
	output, err := g.Runner.Run(orcaConfigDir, "git", "push")
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}

func (g *GitCLIClient) PushSetUpstream(orcaConfigDir, remote, branch string) (string, error) {
	output, err := g.Runner.Run(orcaConfigDir, "git", "push", "-u", remote, branch)
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}
