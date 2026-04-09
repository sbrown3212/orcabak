package git

func (g *GitCLIClient) Add(orcaProfileDir string, args ...string) error {
	args = append([]string{"add"}, args...)

	_, err := g.Runner.Run(orcaProfileDir, "git", args...)
	if err != nil {
		return err
	}

	return nil
}
