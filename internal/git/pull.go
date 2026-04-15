package git

func (g *GitCLIClient) PullFastForward(orcaProfileDir string) (string, error) {
	output, err := g.Runner.Run(orcaProfileDir, "git", "pull", "--ff-only")

	return string(output), err
}
