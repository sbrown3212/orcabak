package git

func (g *GitCLIClient) Add(orcaProfileDir string, args ...string) (string, error) {
	args = append([]string{"add"}, args...)

	output, err := g.Runner.Run(orcaProfileDir, "git", args...)
	if err != nil {
		return string(output), err
	}

	return string(output), nil
}
