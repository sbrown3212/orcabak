package git

func (g *GitCLIClient) Add(repoDir string, args ...string) error {
	args = append([]string{"add"}, args...)

	_, err := g.Runner.Run(repoDir, "git", args...)
	if err != nil {
		return err
	}

	return nil
}
