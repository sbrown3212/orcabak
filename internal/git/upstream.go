package git

func (g *GitCLIClient) GetUpstream(orcaProfileDir string) (string, error) {
	upstream, err := g.Runner.Run(
		orcaProfileDir, "git", "rev-parse", "--abbrev-ref", "--symbolic-full-name", "@{u}",
	)
	if err != nil {
		return string(upstream), err
	}

	return string(upstream), nil
}
